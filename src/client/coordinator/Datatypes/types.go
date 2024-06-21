package datatypes

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jellydator/ttlcache/v3"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
	"github.com/witnesschain-com/diligencewatchtower-client/watcher"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

type WSRequestData struct {
	Api             string
	MessageType     string
	RequestId       string
	ChainID         string
	TransactionHash string
}

type WSTracerResponse struct {
	RequestID       string `json:"requestId"`
	Api             string `json:"api"`
	ChainId         string `json:"chainId"`
	TransactionHash string `json:"transactionHash"`
	Result          string `json:"result"`
	Signature       string `json:"signature"`
}

func (s *WSTracerResponse) Size() int {
	jsonMarshallBytes, err := json.Marshal(s)
	if err != nil {
		wtCommon.Error(fmt.Sprintf("Unable to marshall coordinator result: %v", err))
	}
	size := len(jsonMarshallBytes)
	return size
}

type TraceTxnResult struct {
	Receipt opchain.TransactionReceipt
	Status  string
}

type TracerDependencies struct {
	Cache             *ttlcache.Cache[string, watcher.InclusionProof]
	Config            wtCommon.SimplifiedConfig
	WatchtowerAddress common.Address
	Vault             *keystore.Vault
}
