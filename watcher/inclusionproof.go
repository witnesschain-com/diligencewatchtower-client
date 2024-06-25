package watcher

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/jellydator/ttlcache/v3"

	"github.com/witnesschain-com/diligencewatchtower-client/bindings"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/contractutils"
	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
)

// starts the txn inclusion watcher component of the watchtower
func StartInclusionWatcher(
	wg *sync.WaitGroup,
	config *wtCommon.SimplifiedConfig,
) {

	defer wg.Done()

	// creates wait group
	var waitGroup sync.WaitGroup

	// get a persistent websocket connection to Submission chain node
	SubmissionChainClient, err := wtCommon.GetConnection(
		config.ProofSubmissionWebsocketURL,
		config.Retries,
	)
	if err != nil {
		wtCommon.Error(err)
	}
	defer SubmissionChainClient.Close()


	vault, err := keystore.SetupVault(config)

	if err != nil {
		wtCommon.Error(err)
	}

	var inclusionProofBatch InclusionProofBatch

	// get inclusion proof contract instance
	inclusionProofManagerContract, err := bindings.NewDiligenceProofManager(config.DiligenceProofManagerAddress, SubmissionChainClient)
	if err != nil {
		wtCommon.Error(err)
	}

	inclusionProofBatchSize, err := inclusionProofManagerContract.GetRangeForChainID(nil, big.NewInt(config.CurrentL2ChainID))
	if err != nil {
		wtCommon.Error(err)
	}
	inclusionProofBatch.batchSize = inclusionProofBatchSize.Int64()

	// subscribes to L2 chain head (new blocks)
	l2_sub_head, headers, err := opchain.ListenForHead(config)
	if err != nil {
		wtCommon.Error(err)
	}

	// initialize cache
	cache := ttlcache.New[string, InclusionProof](
		ttlcache.WithTTL[string, InclusionProof](2 * time.Minute),
	)
	go cache.Start() // starts automatic expired item deletion

	for {
		select {
		// if there's an error in subscription, log it
		case err := <-l2_sub_head.Err():
			{
				wtCommon.Error(err)
			}
			// new block notification recieved, generate PoI
		case newBlock := <-headers:
			{
				wtCommon.StartBlock("New L2 block found")
				wtCommon.Success(newBlock.Number.String())

				// fetch receipts for all tnx in the block
				blockReceipts := FetchBlockTxnReceipts(config, newBlock.Hash().Hex())

				// signs each txn receipt and caches them
				_ = SignAndCacheReceiptBatch(blockReceipts, config.WatchtowerAddress, vault, cache)

				// get receipt root hash for signing block PoI
				wtCommon.Info("Getting receipt trie root hash ...")
				wtCommon.Success(newBlock.ReceiptHash)

				inclusionProofBatch.append(InclusionProofHash{proofHash: newBlock.ReceiptHash.Hex()})
				if newBlock.Number.Int64()%inclusionProofBatch.batchSize == 0 {
					if inclusionProofBatch.checkBatchCompletion() {
						// construct the merkle root hash for inclusion proofs and submit

						inclusionProofMerkleRoot := inclusionProofBatch.getRootHash()
						inclusionProofSubmissionStructure := GeneratePoISubmissionBatchInclusionProof(newBlock.Number, inclusionProofMerkleRoot)
						signedInclusionProofSubmissionStructure := SignProofOfInclusion(inclusionProofSubmissionStructure, config.WatchtowerAddress, vault)
						auth, err := contractutils.GetTransactOpts(config, SubmissionChainClient, vault)
						if err != nil {
							wtCommon.Error(err)
						}

						// waitgroup waits for 1 process
						waitGroup.Add(1)
						go contractutils.SubmitProofToInclusionProofManager(&waitGroup, auth, config, newBlock.Number, inclusionProofSubmissionStructure, signedInclusionProofSubmissionStructure, SubmissionChainClient)
						if err != nil {
							wtCommon.Error(err)
						}
					}

					// reset the batch for next batch irrespective of the submission
					// skips partial submission of batch during startup
					inclusionProofBatch.resetBatch()
				}

				// generate the PoI at block level
				PoI := GenerateBlockInclusionProof(
					newBlock.Number,
					newBlock.ReceiptHash,
				)

				// sign the block level PoI
				_ = SignProofOfInclusion(PoI, config.WatchtowerAddress, vault)

				wtCommon.EndBlock("New L2 block found")

			}
		}
	}

}

// Generates and returns the PoI byte array
// for a given block number and receipt root hash
// params:
//  1. L2 block number: big.Int
//  2. Receipt root hash: "github.com/ethereum/go-ethereum/common".Hash
func GenerateBlockInclusionProof(
	blockNumber *big.Int,
	receiptRootHash ethCommon.Hash,
) []byte {

	version_number := big.NewInt(0)
	wtCommon.Info("the `version_number` for proof of inclusion is")
	wtCommon.Success(version_number)
	wtCommon.Info("Computing proof of inclusion")
	proofOfInclusion := []byte(blockNumber.String() + "," + receiptRootHash.Hex() + "," + version_number.String())
	return proofOfInclusion

}

func GeneratePoISubmissionBatchInclusionProof(
	blockNumber *big.Int,
	inclusionProofMerkleRoot []byte,
) []byte {
	version_number := big.NewInt(0)
	wtCommon.Info("the `version_number` for proof of inclusion is")
	wtCommon.Success(version_number)
	wtCommon.Info("Computing proof of inclusion for current batch")
	proofOfInclusion := []byte(blockNumber.String() + "," + hex.EncodeToString(inclusionProofMerkleRoot) + "," + version_number.String())
	return proofOfInclusion
}

// fetches the txn receipts of all the txns in a given block from the L2 node
// params:
//  1. block hash: string
func FetchBlockTxnReceipts(
	config *wtCommon.SimplifiedConfig,
	blockHash string,
) []opchain.TransactionReceipt {
	wtCommon.Info("Fetching receipts for all txn in block with Hash")
	wtCommon.Success(blockHash)
	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// interface for result of the block receipts
	// `receipts` will hold the list of receipts of txns in the block
	var receipts []opchain.TransactionReceipt

	// make the call and collect receipts
	err = client.Call(&receipts, "eth_getBlockReceipts", blockHash)
	if err != nil {
		wtCommon.Fatal(err)
	}
	wtCommon.Success("Receipts fetched successfully!")

	// for i := 0; i < len(receipts); i++ {
	// 	wtCommon.Success(fmt.Sprintf("%+v", receipts[i]))
	// }

	return receipts
}

// computes and signs the proof of inclusion of the given txn batch
// params:
// 1. list of txn hashes: string[]
// 2. pointer to simplified config object: *wtCommon.SimplifiedConfig
func GetProofOfInclusion(
	txnHashBatch []string,
	config *wtCommon.SimplifiedConfig,
	vault *keystore.Vault,
	cache *ttlcache.Cache[string, InclusionProof],

) []InclusionProof {

	var queryBatch []string
	var queryMapping []int
	var inclusionProofs = make([]InclusionProof, len(txnHashBatch))

	for i, txnHash := range txnHashBatch {
		if cache.Has(txnHash) {
			inclusionProofs[i] = cache.Get(txnHash).Value()
		} else {
			queryBatch = append(queryBatch, txnHash)
			queryMapping = append(queryMapping, i)
		}
	}

	batchReceipts := opchain.FetchTxnBatchReceipts(queryBatch, config)
	signedReceipts := SignAndCacheReceiptBatch(batchReceipts, config.WatchtowerAddress, vault ,nil)

	for i, POI := range signedReceipts {
		inclusionProofs[queryMapping[i]] = POI
	}

	return inclusionProofs
}

func SignProofOfInclusion(
	proofOfInclusion []byte,
	watchtower common.Address,
	vault *keystore.Vault,
) []byte {

	hash := crypto.Keccak256Hash(proofOfInclusion)

	wtCommon.Info("Signing Proof of Inclusion")


	signature, err := vault.SignData(hash.Bytes())
	if err != nil {
		wtCommon.Error(err)
	}

	return signature
}

func SignAndCacheReceiptBatch(
	receipts []opchain.TransactionReceipt,
	watchtower common.Address,
	vault *keystore.Vault,
	cache *ttlcache.Cache[string, InclusionProof],
) []InclusionProof {

	var signedReceipts []InclusionProof

	for index := 0; index < len(receipts); index++ {
		marshalledReceipt, err := receipts[index].MarshalJSON()
		if err != nil {
			wtCommon.Error("Failed to marshall receipt!")
			wtCommon.Warning(fmt.Sprintf("%+v", receipts[index]))
			marshalledReceipt = []byte("")
		}
		signedReceipt := SignProofOfInclusion(marshalledReceipt, watchtower, vault)
		SignedPoI := InclusionProof{receipts[index], signedReceipt}
		if cache != nil {
			cache.Set(receipts[index].TxHash.Hex(), SignedPoI, ttlcache.DefaultTTL)
		}
		signedReceipts = append(signedReceipts, SignedPoI)
	}
	return signedReceipts
}
