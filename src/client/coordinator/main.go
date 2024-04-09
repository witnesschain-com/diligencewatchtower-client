package coordinator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	wtCommon "github.com/diligencewatchtower-client/common"
	datatypes "github.com/diligencewatchtower-client/coordinator/Datatypes"
	auth "github.com/diligencewatchtower-client/coordinator/auth"
	ws "github.com/diligencewatchtower-client/coordinator/core"
	wtInterface "github.com/diligencewatchtower-client/coordinator/interfaces"
	watcher "github.com/diligencewatchtower-client/watcher"
	"github.com/gorilla/websocket"
	"github.com/jellydator/ttlcache/v3"
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

func getWatchingChainID(cfg wtCommon.WatchTowerConfig) string {
	watchingL2 := cfg.CurrentlyWatchingL2
	chainList := cfg.L2.OPChains
	for _, chain := range chainList {
		if chain.Name == watchingL2 {
			return strconv.Itoa(chain.ChainID)
		}
	}
	return ""
}

func StartCoordinator(config wtCommon.WatchTowerConfig) {
	client := auth.CoordinatorClient{}
	if chainId := getWatchingChainID(config); chainId != "" {
		err := client.Initialize(config.WitnesschainCoordinatorUrl, config.PrivateKey, chainId)
		if err != nil {
			wtCommon.Error(err)
			return
		}
	} else {
		wtCommon.Error("ChainID not found for " + config.CurrentlyWatchingL2)
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
	wsHandler := ws.WebsockerClient{}

	err = wsHandler.ConnectToCoordinator(headers)
	if err != nil {
		wtCommon.Error(err)
		return
	}

	cache := ttlcache.New(
		ttlcache.WithTTL[string, watcher.InclusionProof](CACHE_EXPIRY),
	)
	go cache.Start()

	var simplifiedConfig wtCommon.SimplifiedConfig
	wtCommon.LoadSimplifiedConfig(&config, &simplifiedConfig)

	dependencies := datatypes.TracerDependencies{
		Cache:      cache,
		Config:     simplifiedConfig,
		PrivateKey: config.PrivateKey,
	}

	var DataChannel = make(chan string, 5000)

	go HandleMessages(wsHandler.Connection, DataChannel, dependencies)
	go Heartbeat(wsHandler.Connection)
	err = wsHandler.ListenForMessages(DataChannel)
	wtCommon.Error(err)
	wsHandler.CloseConnection()

}

func HandleMessages(connection *websocket.Conn, channel chan string, deps datatypes.TracerDependencies) {
	var lastTxn string
	var lastChainID string
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
				signature, _ := auth.SignCoordinatorMessage(string(payloadBytes), deps.PrivateKey)

				coordinatorResp := datatypes.WSTracerResponse{
					RequestID:       wsRequest.RequestId,
					Api:             wsRequest.Api,
					ChainId:         wsRequest.ChainID,
					TransactionHash: wsRequest.TransactionHash,
					Result:          string(payloadBytes),
					Signature:       signature,
				}

				logData, err := json.Marshal(coordinatorResp)
				if err != nil {
					wtCommon.Error(err)
				} else {
					cutoff := len(string(logData))
					if cutoff > 50 {
						cutoff = 50
					}
					wtCommon.Info(fmt.Sprintf("WS:SEN Type[%d] Content[%s]", websocket.TextMessage, string(logData)))

					err = connection.WriteJSON(coordinatorResp)
					if err != nil {
						wtCommon.Error(err)
						break
					}
				}
			}
		}

	}
	wtCommon.Warning("Terminating message handler")
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

	return &datatypes.WSRequestData{
		Api:             m.(map[string]interface{})["api"].(string),
		MessageType:     m.(map[string]interface{})["messageType"].(string),
		RequestId:       m.(map[string]interface{})["requestId"].(string),
		ChainID:         m.(map[string]interface{})["chainId"].(string),
		TransactionHash: m.(map[string]interface{})["transactionHash"].(string),
	}, nil
}
