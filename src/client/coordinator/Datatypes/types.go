package datatypes

import (
	"crypto/ecdsa"

	"github.com/diligencewatchtower-client/opchain"
	"github.com/diligencewatchtower-client/watcher"
	"github.com/jellydator/ttlcache/v3"

	wtCommon "github.com/diligencewatchtower-client/common"
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

type TraceTxnResult struct {
	Receipt opchain.TransactionReceipt
	Status  string
}

type TracerDependencies struct {
	Cache      *ttlcache.Cache[string, watcher.InclusionProof]
	Config     wtCommon.SimplifiedConfig
	PrivateKey *ecdsa.PrivateKey
}
