package coordinator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jellydator/ttlcache/v3"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	datatypes "github.com/witnesschain-com/diligencewatchtower-client/coordinator/Datatypes"
	auth "github.com/witnesschain-com/diligencewatchtower-client/coordinator/auth"
	"github.com/witnesschain-com/diligencewatchtower-client/coordinator/core"
	ws "github.com/witnesschain-com/diligencewatchtower-client/coordinator/core"
	wtInterface "github.com/witnesschain-com/diligencewatchtower-client/coordinator/interfaces"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	watcher "github.com/witnesschain-com/diligencewatchtower-client/watcher"
)

type Counter struct {
	mu                sync.Mutex
	PingsWithoutPongs int
}

const (
	HEARTBEAT_INTERVAL = 30 * time.Second
)

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.PingsWithoutPongs += 1
}

func (c *Counter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.PingsWithoutPongs = 0
}

const (
	HEARTBEAT_TIMEOUT = 5
	CACHE_EXPIRY      = 2 * time.Minute
)

func getWatchingChainID(simpleConfig wtCommon.SimplifiedConfig) string {
	if simpleConfig.CurrentL2ChainID != 0 {
		return strconv.Itoa(int(simpleConfig.CurrentL2ChainID))
	}
	return ""
}

func StartCoordinator(simpleConfig wtCommon.SimplifiedConfig) {
	client := auth.CoordinatorClient{}
	if chainId := getWatchingChainID(simpleConfig); chainId != "" {
		err := client.Initialize(simpleConfig.WitnesschainCoordinatorURL, simpleConfig.WatchtowerAddress, chainId, &simpleConfig)
		if err != nil {
			wtCommon.Error(err)
			return
		}
	} else {
		wtCommon.Error("ChainID not found for " + simpleConfig.CurrentL2Chain)
		return
	}

	ok, err := client.AuthenticateToWitnesschain()
	if err != nil {
		// commenting the error for beta release as this is not intended to be part of launch
		// and causes unnecessary error logs. This will be uncommendted in the next stable version
		// wtCommon.Error(err)
		return
	}
	if !ok {
		return
	}

	headers := client.GetHeaders()
	wsHandler := ws.WebsocketClient{}

	err = wsHandler.ConnectToCoordinator(headers)
	if err != nil {
		wtCommon.Error(err)
		return
	}

	cache := ttlcache.New(
		ttlcache.WithTTL[string, watcher.InclusionProof](CACHE_EXPIRY),
	)
	go cache.Start()

	Vault, err := keystore.SetupVault(&simpleConfig)
	if err != nil {
		wtCommon.Error(err)
	}

	dependencies := datatypes.TracerDependencies{
		Cache:             cache,
		Config:            simpleConfig,
		WatchtowerAddress: simpleConfig.WatchtowerAddress,
		Vault:             Vault,
	}

	var DataChannel = make(chan string, 5000)

	go HandleMessages(&wsHandler, DataChannel, dependencies, &simpleConfig, client)
	go Heartbeat(wsHandler.Connection)
	err = wsHandler.ListenForMessages(DataChannel)
	wtCommon.Error(err)
	wsHandler.CloseConnection()

}

func HandleMessages(ws *core.WebsocketClient, channel chan string, deps datatypes.TracerDependencies, config *wtCommon.SimplifiedConfig, client auth.CoordinatorClient) {
	var lastTxn string
	var lastChainID string
	signer, err := keystore.SetupVault(config)
	if err != nil {
		wtCommon.Error(err)
	}

	connection := ws.Connection
	for {
		data := <-channel
		wsRequest, err := handleWsJson(data)
		if err != nil {
			wtCommon.Error(err)
			continue
		}
		if lastTxn == wsRequest.TransactionHash && lastChainID == wsRequest.ChainID {
			wtCommon.Info("Ignoring Already served tracer-request")
		} else {
			lastChainID = wsRequest.ChainID
			lastTxn = wsRequest.TransactionHash
			if wsRequest.MessageType == "request" && wsRequest.Api == "trace-transaction" {
				res := wtInterface.TraceTransaction(wsRequest.TransactionHash, deps)
				payloadBytes, _ := json.Marshal(res)
				signature, _ := auth.SignCoordinatorMessage(string(payloadBytes), deps.WatchtowerAddress, signer)

				coordinatorResp := datatypes.WSTracerResponse{
					RequestID:       wsRequest.RequestId,
					Api:             wsRequest.Api,
					ChainId:         wsRequest.ChainID,
					TransactionHash: wsRequest.TransactionHash,
					Result:          string(payloadBytes),
					Signature:       signature,
				}
				restApiData, err := json.Marshal(coordinatorResp)
				if err != nil {
					wtCommon.Error(fmt.Sprintf("Unable to marshall coordinator response: %v", err))
				}

				wsSuccessful := true
				wtCommon.Info("WS: Submitting txn-tracer results..")
				if coordinatorResp.Size() < ws.GetWriteBufferSize() {
					if err := handleWebsocketResultSubmission(connection, coordinatorResp); err != nil {
						wtCommon.Error(err)
						wsSuccessful = false
						break
					}

				}
				if !wsSuccessful {
					// 5 retries
					for i := 0; i < 5; i++ {
						if err := handleRestApiSubmission(client, restApiData); err == nil {
							break
						}
					}
				}
			}
		}

	}
	wtCommon.Warning("Terminating message handler")
}

func handleWebsocketResultSubmission(connection *websocket.Conn, coordinatorResp datatypes.WSTracerResponse) error {
	err := connection.WriteJSON(coordinatorResp)
	if err != nil {
		return err
	}
	return nil
}

func handleRestApiSubmission(client auth.CoordinatorClient, data []byte) error {
	statusCode, err := client.SubmitResult(data)
	if statusCode == 200 {
		wtCommon.Success("Submitted result to coordinator")
	}
	if err != nil {
		wtCommon.Error(err)
		return err
	}
	return nil
}

func Heartbeat(connection *websocket.Conn) {
	wtCommon.Info("Starting heartbeat routine...")
	defer func() {
		wtCommon.Warning("Closing heartbeat routine")
		connection.Close()
	}()
	for {
		err := connection.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			wtCommon.Error(err)
			return
		}
		// wtCommon.Info(fmt.Sprintf("WS:SEN Type[%d] Content[%s]", websocket.TextMessage, "ping"))
		time.Sleep(HEARTBEAT_INTERVAL)
	}
}

func handleWsJson(data string) (*datatypes.WSRequestData, error) {
	bytes := []byte(data)
	var m interface{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}

	if m.(map[string]interface{})["messageType"].(string) == "error" {
		wtCommon.Error(m.(map[string]interface{})["error"].(map[string]interface{})["message"].(string))
	}

	return &datatypes.WSRequestData{
		Api:             m.(map[string]interface{})["api"].(string),
		MessageType:     m.(map[string]interface{})["messageType"].(string),
		RequestId:       m.(map[string]interface{})["requestId"].(string),
		ChainID:         m.(map[string]interface{})["chainId"].(string),
		TransactionHash: m.(map[string]interface{})["transactionHash"].(string),
	}, nil
}
