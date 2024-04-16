package contractutils

import (
	"strconv"

	"math/big"

	"github.com/witnesschain-com/diligencewatchtower-client/bindings"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func SubmitProofToAlertManager(
	auth *bind.TransactOpts,
	config *wtCommon.SimplifiedConfig,
	blockNumber *big.Int,
	outputRootL2OO string,
	outputRootWatchtower string,
	proofOfDilegence []byte,
) (*types.Transaction, error) {

	// get a persistent websocket connection to L1 chain node
	L1Client, err := wtCommon.GetConnection(
		config.ProofSubmissionWebsocketURL,
		config.Retries,
	)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer L1Client.Close()

	// create the AlertManager contract object from address
	alertManagerContract, err := bindings.NewAlertManager(config.AlertManagerAddress, L1Client)
	if err != nil {
		wtCommon.Fatal(err)
	}

	wtCommon.Info("Calling `RaiseAlert` contract on L1 ...")

	outputRootL2OOBytes := []byte(outputRootL2OO)
	outputRootWatchtowerBytes := []byte(outputRootWatchtower)

	txn, err := alertManagerContract.RaiseAlert(
		auth,
		big.NewInt(config.CurrentL2ChainID),
		blockNumber,
		outputRootL2OOBytes,
		outputRootWatchtowerBytes,
		proofOfDilegence,
	)

	if err != nil {
		for retryCount := 0; retryCount < config.Retries; retryCount++ {
			wtCommon.Error(err)
			wtCommon.Warning("Failed to raise alert, retrying, attempt :" + strconv.Itoa(retryCount))
			txn, err = alertManagerContract.RaiseAlert(
				auth,
				big.NewInt(config.CurrentL2ChainID),
				blockNumber,
				outputRootL2OOBytes,
				outputRootWatchtowerBytes,
				proofOfDilegence,
			)
			if err == nil {
				break
			}
		}
		if err != nil {
			wtCommon.Error(err)
			wtCommon.Fatal("Failed despite " + strconv.Itoa(config.Retries) + " retries, please add funds if you lack gas or report the error to the admin")
		}
	}
	wtCommon.Success("Alert raised successfully with transaction hash ...")
	wtCommon.Success(txn.Hash().Hex())

	return txn, nil
}
