/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package watcher

import (
	"strconv"
	"sync"

	// "github.com/ethereum/go-ethereum/accounts"
	"github.com/witnesschain-com/diligencewatchtower-client/L1"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/contractutils"
	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
)

var parsed_output *opchain.OutputProposed = nil
var globalConfigData *wtCommon.SimplifiedConfig = nil

// High level abstract function which does the following:
// 1. Validates the state commitment made by rollup on L1
// 2. Raises alert on `AlertManager` contract if found invalid
// 3. Computes and submits ProofOfDiligence on `DiligenceProofManager` or `AlertManager`
// params:
//  1. `L2OO_output`	: parsed event log from L2 output oracle contract
func do_tracing(
	L2OO_output *opchain.OutputProposed,
	vault *keystore.Vault,
) bool {

	// get a persistent websocket connection to Submission chain node
	SubmissionChainClient, err := wtCommon.GetConnection(
		globalConfigData.ProofSubmissionWebsocketURL,
		globalConfigData.Retries,
	)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer SubmissionChainClient.Close()

	// fetch block hashes for l2 block at proposed output and it's parent
	// also get output_root from L2 node for validation against L2OO
	ultimateBlockHash, penultimateBlockHash, L2OutputRoot := ProcessIntermediateBlocks(L2OO_output.L2BlockNumber, globalConfigData)

	// validate output assertion and set flag to store result of assertion validation
	var isValid bool = false
	if opchain.ValidateOutputRoot(L2OutputRoot, L2OO_output.OutputRoot) {
		isValid = true
	}

	// run `debug_intermediateRoots` tracer to get intermediate state roots
	// for ultimate and penultimate blocks to be used in PoD computation
	midPointPenultimateBlock, midPointUltimate := FetchIntermediateStateRoots(
		globalConfigData,
		ultimateBlockHash,
		penultimateBlockHash,
	)

	// compute proof of diligence
	proofOfDilegence := ComputeProofOfDiligence(
		L2OO_output.L2BlockNumber,
		midPointPenultimateBlock,
		midPointUltimate,
	)

	// sign proof of diligence and get the final Ethereum Signed Message for it
	signatureOfProofOfDiligence := SignProofOfDiligence(
		proofOfDilegence,
		globalConfigData.WatchtowerAddress,
		vault,
	)

	auth, err := contractutils.GetTransactOpts(globalConfigData, SubmissionChainClient, vault)
	if err != nil {
		wtCommon.Error(err)
	}

	if isValid {
		// submit proof of diligence to `DiligenceProofManager` if assertion was found valid
		txn, err := contractutils.SubmitProofToDiligenceProofManager(
			auth,
			globalConfigData,
			L2OO_output.L2BlockNumber,
			proofOfDilegence,
			signatureOfProofOfDiligence,
			SubmissionChainClient,
		)
		if err != nil {
			wtCommon.Error(err)
		}
		// wait for txn receipt to ensure successful submission
		_, err = contractutils.WaitForTransactionReceipt(
			globalConfigData,
			txn,
			SubmissionChainClient,
		)
		if err != nil {
			wtCommon.Error(err)
		}
	} else {
		// Raise alert and submit proof of diligence to `AlertManager`
		// if assertion was found to be invalid
		txn, err := contractutils.SubmitProofToAlertManager(
			auth,
			globalConfigData,
			L2OO_output.L2BlockNumber,
			L2OO_output.OutputRoot,
			L2OutputRoot,
			proofOfDilegence,
		)
		if err != nil {
			wtCommon.Error(err)
		}

		// wait for txn receipt to ensure successful submission
		_, err = contractutils.WaitForTransactionReceipt(
			globalConfigData,
			txn,
			SubmissionChainClient,
		)
		if err != nil {
			wtCommon.Error(err)
		}
	}

	// end block and return to go back on waiting for next proposal
	wtCommon.EndBlock("New L2 Block proposed")
	return true
}

// Subscribes to a L2OO's event logs emmitted on the node (L1)
// params: waitgroup for go-routine synchronization and pointer to config
func StartWatcher(
	wg *sync.WaitGroup,
	configData *wtCommon.WatchTowerConfig,
	configChan chan bool,
) bool {
	defer wg.Done()

	wtCommon.Info("Determining chain assignment ...")
	globalConfigData = wtCommon.LoadSimplifiedConfig(configData, globalConfigData)
	wtCommon.Success("Rollup chain assigned is: `" + configData.CurrentlyWatchingL2 + "`")

	// subscribe to state commitment events on L1
	_, eth_sub, eth_logs, err := L1.ListenForProposals(globalConfigData)
	if err != nil {
		wtCommon.Fatal(err)
	}

	signer, err := 	keystore.SetupVault(globalConfigData)

	if err != nil {
		wtCommon.Error(err)
	}

	number_of_retries := 0

	go StartInclusionWatcher(wg, globalConfigData)
	for {

		if len(configChan) != 0 {

			restart := <-configChan

			wtCommon.Info("WatchTower configuration has been updated; applying the updates ...")
			configData = wtCommon.LoadConfigFromJson()
			wtCommon.Success("Reloaded the new configuration from `wtCommon.json`")

			globalConfigData = wtCommon.LoadSimplifiedConfig(configData, globalConfigData)
			if restart {
				eth_sub.Unsubscribe()
				close(eth_logs)
				_, eth_sub, eth_logs, err = L1.ListenForProposals(globalConfigData)
				if err != nil {
					wtCommon.Fatal(err)
				}
			}
		}

		if parsed_output != nil {

			// there was something to be processed but we crashed
			// while processing it

			for number_of_retries < configData.Retries {

				if do_tracing(parsed_output, signer) == true {
					number_of_retries = 0
					parsed_output = nil
					break
				} else {
					number_of_retries++
				}
			}

			if number_of_retries >= configData.Retries {

				wtCommon.Warning("Please check if you are out of gas")
				wtCommon.Fatal("Failed after `" + strconv.Itoa(configData.Retries) + "` retries")
			}
		}

		select {

		case err := <-eth_sub.Err():
			{
				wtCommon.Error(err)
			}

		case filteredLog := <-eth_logs:
			{
				wtCommon.StartBlock("New L2 block proposed")

				parsed_output = opchain.ParseOutputProposed(filteredLog)

				if do_tracing(parsed_output, signer) == true {
					wtCommon.Info("Waiting for next proposal ...")
					parsed_output = nil
				}
			}
		}
	}

	return true
}
