package contractutils

import (
	"strconv"
	"sync"

	"math/big"

	"github.com/diligencewatchtower-client/bindings"
	wtCommon "github.com/diligencewatchtower-client/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubmitProofToDiligenceProofManager(
	auth *bind.TransactOpts,
	config *wtCommon.SimplifiedConfig,
	blockNumber *big.Int,
	proofOfDilegence []byte,
	signatureOfProofOfDiligence []byte,
	SubmissionChainClient *ethclient.Client,
) (*types.Transaction, error) {

	// create the DiligenceProofManager contract object from address
	diligenceProofManagerContract, err := bindings.NewDiligenceProofManager(config.DiligenceProofManagerAddress, SubmissionChainClient)
	if err != nil {
		wtCommon.Fatal(err)
	}
	wtCommon.Info("Calling `SubmitProof` contract on L1 ...")

	txn, err := diligenceProofManagerContract.SubmitPODProof(
		auth,
		big.NewInt(config.CurrentL2ChainID),
		blockNumber,
		proofOfDilegence,
		signatureOfProofOfDiligence,
	)

	if err != nil {
		for retryCount := 0; retryCount < config.Retries; retryCount++ {
			wtCommon.Error(err)
			wtCommon.Warning("Failed to submit proof of diligence, retrying, attempt :" + strconv.Itoa(retryCount))
			txn, err = diligenceProofManagerContract.SubmitPODProof(
				auth,
				big.NewInt(config.CurrentL2ChainID),
				blockNumber,
				proofOfDilegence,
				signatureOfProofOfDiligence,
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
	wtCommon.Success("Bounty successfully mined and submitted with transaction hash ...")
	wtCommon.Success(txn.Hash().Hex())
	return txn, nil
}

func SubmitProofToInclusionProofManager(
	wg *sync.WaitGroup,
	auth *bind.TransactOpts,
	config *wtCommon.SimplifiedConfig,
	blockNumber *big.Int,
	proofOfInclusionRootHash []byte,
	signatureOfProofOfInclusionRootHash []byte,
	SubmissionChainClient *ethclient.Client,
) {
	defer wg.Done()

	// create the InclusionProofManager contract object from address
	inclusionProofManagerContract, err := bindings.NewDiligenceProofManager(config.DiligenceProofManagerAddress, SubmissionChainClient)
	if err != nil {
		wtCommon.Fatal(err)
	}
	wtCommon.Info("Submitting Proof of Inclusion to submission chain...")

	txn, err := inclusionProofManagerContract.SubmitPOIProof(
		auth,
		big.NewInt(config.CurrentL2ChainID),
		blockNumber,
		proofOfInclusionRootHash,
		signatureOfProofOfInclusionRootHash,
	)

	if err != nil {
		for retryCount := 0; retryCount < config.Retries; retryCount++ {
			wtCommon.Error(err)
			wtCommon.Warning("Failed to submit proof of inclusion, retrying, attempt :" + strconv.Itoa(retryCount))
			txn, err = inclusionProofManagerContract.SubmitPOIProof(
				auth,
				big.NewInt(config.CurrentL2ChainID),
				blockNumber,
				proofOfInclusionRootHash,
				signatureOfProofOfInclusionRootHash,
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
	wtCommon.Success("PoI Bounty successfully mined and submitted with transaction hash ...")
	wtCommon.Success(txn.Hash().Hex())

	// wait for txn receipt to ensure successful submission
	_, err = WaitForTransactionReceipt(
		config,
		txn,
		SubmissionChainClient,
	)
	if err != nil {
		wtCommon.Error(err)
	}
}
