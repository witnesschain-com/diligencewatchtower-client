/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package opchain

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetHeaderByHash(blockHash string, config *wtCommon.SimplifiedConfig) map[string]string {

	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// `result` will hold the header of the block
	var result map[string]string

	// make the eth call and collect result
	err = client.Call(&result, "eth_getHeaderByHash", blockHash)
	if err != nil {
		wtCommon.Error(err)
		return nil
	}

	return result

}

func GetHeaderByBlockNumber(blockNumber string, config *wtCommon.SimplifiedConfig) map[string]string {

	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// `result` will hold the header of the block
	var result map[string]string

	// make the eth call and collect result
	err = client.Call(&result, "eth_getBlockByNumber", blockNumber)
	if err != nil {
		wtCommon.Error(err)
		return nil
	}

	return result

}

func GetTxnByHash(txnHash string, config *wtCommon.SimplifiedConfig) map[string]string {

	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// `result` will hold the header of the block
	var result map[string]string

	// make the eth call and collect result
	err = client.Call(&result, "eth_getTransactionByHash", txnHash)
	if err != nil {
		wtCommon.Error(err)
		return nil
	}

	return result

}

func GetOutputRootByBlockNumber(blockNumber string, nodeRPC_URL string) string {

	data, _ := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "optimism_outputAtBlock",
		"params":  []string{blockNumber},
		"id":      1,
	})

	requestBody := bytes.NewBuffer(data)
	numFailures := -1

get_output:

	numFailures++
	if numFailures > 5 {
		wtCommon.Fatal("Failed to fetch output from L2 node")
	}

	resp, err := http.Post(nodeRPC_URL, "application/json", requestBody)
	if err != nil {
		wtCommon.Error(err)
		goto get_output
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wtCommon.Error(err)
		goto get_output
	}

	var respData OptimismNodeOutputStruct
	if err := json.Unmarshal(respBody, &respData); err != nil {
		// wtCommon.Warning("Can not unmarshal JSON")
		wtCommon.Warning(fmt.Sprintf("Resp code: <%v>, Cannot unmarshal JSON : code: <%v>", resp.StatusCode, respBody))
		wtCommon.Info("Awaiting the `output_root` on L2 node, sleeping for a sec")
		time.Sleep(1 * time.Second)
		wtCommon.Error(err)
		goto get_output

	}

	if respData.Result.OutputRoot == "" {
		wtCommon.Info("Awaiting the `output_root` on L2 node, sleeping for a sec")
		time.Sleep(1 * time.Second)
		goto get_output
	}

	return respData.Result.OutputRoot

}

// fetches the txn receipt of a given txn from the L2 node
// params:
//  1. tnx hash: string
func FetchTxnReceipt(txnHash string, config *wtCommon.SimplifiedConfig) TransactionReceipt {
	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// interface for result of the transaction receipt
	// `receipt` will hold the txn receipt as a struct
	var receipt TransactionReceipt

	// make the eth rpc call and collect receipt
	err = client.Call(&receipt, "eth_getTransactionReceipt", txnHash)
	if err != nil {
		wtCommon.Error(err)
	} else {
		wtCommon.Info("RECEIPT IS : ")
		wtCommon.Success(fmt.Sprintf("%+v\n", receipt))
	}
	wtCommon.Info("   ")
	return receipt
}

// fetches the txn receipt of a given set of txns from the L2 node
// params:
//  1. list of txn hashes: string[]
func FetchTxnBatchReceipts(txnHashBatch []string, config *wtCommon.SimplifiedConfig) []TransactionReceipt {
	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// convert txn hashes to []interface{} type
	txnHashBatchInterfaces := make([]interface{}, len(txnHashBatch))
	for i, v := range txnHashBatch {
		txnHashBatchInterfaces[i] = v
	}

	// declate array of `rpc.BatchElem`` objects for batched requests
	// each entry in array defines a receipt request
	var batch []rpc.BatchElem

	// declare an empty array of errors to collect any errors that occur
	// per individual receipt request in the batch (doesn't include i/o error)
	var errs []rpc.Error

	// array of interface for result of the transaction receipt
	// each element in `receipts` will hold the txn receipt as a struct
	var receipts []TransactionReceipt

	// populate the receipts and errors arrays/slices with empty structs
	// important to do in a separate loop
	// before creating actual request batch element object
	for txnNo := 0; txnNo < len(txnHashBatch); txnNo++ {
		// initialize receipts slice
		receiptResult := new(TransactionReceipt)
		receipts = append(receipts, *receiptResult)

		// initialize errors slice
		reqError := new(rpc.Error)
		errs = append(errs, *reqError)
	}

	// build individual request as a `rpc.BatchElem` and
	// populate the batch slice with them
	for txnNo := 0; txnNo < len(txnHashBatch); txnNo++ {
		receiptRequest := new(rpc.BatchElem)
		receiptRequest.Method = "eth_getTransactionReceipt"
		receiptRequest.Args = []interface{}{txnHashBatchInterfaces[txnNo]}
		receiptRequest.Result = &receipts[txnNo]
		receiptRequest.Error = errs[txnNo]
		batch = append(batch, *receiptRequest)

	}

	// make the eth rpc call and collect receipts
	err = client.BatchCall(batch)
	if err != nil {
		// TODO: remove change FATAL to ERROR, return the error and handle it
		wtCommon.Fatal(err)
	}
	wtCommon.Info("batch fetched receipts are : ")
	wtCommon.Success(fmt.Sprintf("%+v\n", receipts))
	wtCommon.Info("   ")
	return receipts
}

type OptimismSyncStatus struct {
	CurrentL1 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"current_l1"`
	CurrentL1Finalized struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"current_l1_finalized"`
	HeadL1 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"head_l1"`
	SafeL1 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"safe_l1"`
	FinalizedL1 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
	} `json:"finalized_l1"`
	UnsafeL2 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"unsafe_l2"`
	SafeL2 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"safe_l2"`
	FinalizedL2 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"finalized_l2"`
	PendingSafeL2 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"pending_safe_l2"`
	QueuedUnsafeL2 struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"queued_unsafe_l2"`
	EngineSyncTarget struct {
		Hash       string `json:"hash"`
		Number     int    `json:"number"`
		ParentHash string `json:"parentHash"`
		Timestamp  int    `json:"timestamp"`
		L1Origin   struct {
			Hash   string `json:"hash"`
			Number int    `json:"number"`
		} `json:"l1origin"`
		SequenceNumber int `json:"sequenceNumber"`
	} `json:"engine_sync_target"`
}

// Method to determine the current status of a transaction
// params:
//  1. `txnHash` : string : self explainatory
//  2. `config`	 : *wtCommon.SimplifiedConfig : reference to the config
//
// returns:
//
//	Transaction Status : int
//		0 = unsafe (in a block obtainted via p2p but not L1)
//		1 = safe (derived from a recent, unfinalized L1 block)
//		2 = finalized (derived from a finalized L1 block)
func FetchTxnStatus(txnHash string, config *wtCommon.SimplifiedConfig) int {
	// get receipt to know block num of the txn
	Txn := FetchTxnReceipt(txnHash, config)

	// get txn status
	return FetchTxnStatusFromReceipt(Txn, config)
}

// Method to determine the current status of the L2 node
// params:
//  1. `config` : *wtCommon.SimplifiedConfig : reference to the config
//
// returns:
//
//	L2 safe head block number : int
//	L2 finalized head block number : int
func FetchChainStatus(config *wtCommon.SimplifiedConfig) (int, int) {
	wtCommon.Info("Fetching chain status ...")

	// prepare data for the status rpc request
	data, _ := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "optimism_syncStatus",
		"id":      1,
	})

	requestBody := bytes.NewBuffer(data)

	// track failed requests to enable finite retries
	numFailures := -1

get_status:

	numFailures++
	// too many failures, crash
	if numFailures > config.Retries {
		wtCommon.Fatal("Failed to fetch status from L2 node")
	}

	// make the request to the L2 node
	resp, err := http.Post(config.L2NodeRPCURL, "application/json", requestBody)
	if err != nil {
		wtCommon.Error(err)
		goto get_status
	}
	defer resp.Body.Close()

	// read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		wtCommon.Error(err)
		goto get_status
	}

	// interface to hold the unmarshalled response
	var respData map[string]interface{}

	// unmarshall json
	// report and retry if any errors occur
	if err := json.Unmarshal(respBody, &respData); err != nil {
		wtCommon.Warning(fmt.Sprintf("Resp code: <%v>, Cannot unmarshal JSON : code: <%s>", resp.StatusCode, respBody))
		wtCommon.Error(err)
		goto get_status
	}

	// L2 head statuses
	// unsafe = obtained via p2p
	// safe = derived from unfinalized L1
	// finalized = derived from finalized L1
	safe_head := respData["result"].(map[string]interface{})["safe_l2"].(map[string]interface{})["number"].(float64)
	finalized_head := respData["result"].(map[string]interface{})["finalized_l2"].(map[string]interface{})["number"].(float64)

	wtCommon.Success("Safe: " + strconv.Itoa(int(safe_head)) + " Finalized: " + strconv.Itoa(int(finalized_head)))

	return int(safe_head), int(finalized_head)
}

// Method to determine the current status of a transaction
// params:
//  1. `txnReceipt` : TransactionReceipt : the receipt of the txn whose status is being requested
//  2. `config`	 : *wtCommon.SimplifiedConfig : reference to the config
//
// returns:
//
//	Transaction Status : int
//		0 = unsafe (in a block obtainted via p2p but not L1)
//		1 = safe (derived from a recent, unfinalized L1 block)
//		2 = finalized (derived from a finalized L1 block)
func FetchTxnStatusFromReceipt(txnReceipt TransactionReceipt, config *wtCommon.SimplifiedConfig) int {

	wtCommon.Info("Transaction status requested ...")
	// get block num of the txn
	if txnReceipt.BlockNumber == nil {
		wtCommon.Error("Invalid txn receipt, 'BlockNumber' not found!")
		return -1
	}
	cur_block := txnReceipt.BlockNumber.Int64()
	wtCommon.Success("Txn is in block: " + strconv.Itoa(int(cur_block)))

	safe_head, finalized_head := FetchChainStatus(config)

	wtCommon.Info("Determining Txn status ...")

	// L2 head statuses
	// unsafe = obtained via p2p : Txn status = 0
	// safe = derived from unfinalized L1 : Tnx status = 1
	// finalized= derived from finalized L1 : Txn status = 2
	if int(cur_block) <= finalized_head {
		// txn is in a block which is already finalized
		wtCommon.Success("Txn Finalized on L1!")
		return 2
	} else if int(cur_block) <= safe_head {
		// txn is in a block derived from L1 but not finalized
		wtCommon.Success("Txn Derived from L1 but not Finalized!")
		return 1
	}
	// txn was obtained via p2p
	wtCommon.Success("Txn data not posted on L1 yet")
	return 0
}
