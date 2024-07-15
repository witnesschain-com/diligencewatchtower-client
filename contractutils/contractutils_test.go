package contractutils

import (
	"context"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
)


func TestSubmitProofToDiligenceProofManager(t *testing.T){

	config := wtCommon.LoadConfigFromJson("config.json")
	simplifiedConfig := wtCommon.LoadSimplifiedConfig(config, nil)

	vc := keystore.GetVaultConfig(simplifiedConfig)
	vault, err := keystore.SetupVault(vc)
	if err != nil {
		panic(err)
	}

	proofSubmissionClient, err := wtCommon.GetConnection(simplifiedConfig.ProofSubmissionWebsocketURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}

	L2Client, err := wtCommon.GetConnection(simplifiedConfig.L2WebsocketURL, config.Retries)
	if err != nil {
		panic(err)
	}

	blockNumberInt, err := L2Client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	blockNumber := new(big.Int).SetUint64(blockNumberInt) 

	proofOfDilegence := []byte{0x1}
	hash := crypto.Keccak256Hash(proofOfDilegence)
	signedProofOfDiligence, err := vault.SignData(hash.Bytes(), apitypes.TextPlain.Mime)
	if err != nil{
		panic(err)
	}

	auth := vault.NewTransactOpts(big.NewInt(simplifiedConfig.ProofSubmissionChainID))

	wtCommon.Info("watchtower Address: " + auth.From.Hex())
	wtCommon.Info("signature: " + hexutil.Encode(signedProofOfDiligence) )

	transaction, err := SubmitProofToDiligenceProofManager(auth, simplifiedConfig, blockNumber, proofOfDilegence, signedProofOfDiligence, proofSubmissionClient)
	if err != nil {
		t.Fatal(err)
	}
	wtCommon.Info(transaction)
	wtCommon.Info(transaction.Hash())
	recipt, err := WaitForTransactionReceipt(
		simplifiedConfig,
		transaction,
		proofSubmissionClient,
	)
	wtCommon.Info(recipt)
}

func TestSubmitProofToInclusionProofManager(t *testing.T){
	config := wtCommon.LoadConfigFromJson("config.json")
	simplifiedConfig := wtCommon.LoadSimplifiedConfig(config, nil)

	vc := keystore.GetVaultConfig(simplifiedConfig)
	vault, err := keystore.SetupVault(vc)
	if err != nil {
		panic(err)
	}

	proofSubmissionClient, err := wtCommon.GetConnection(simplifiedConfig.ProofSubmissionWebsocketURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}

	auth := vault.NewTransactOpts(big.NewInt(simplifiedConfig.ProofSubmissionChainID))

	// https://sepolia.etherscan.io/tx/0x2c22200a8b55d89cdb780c5fa4ec648ec4421be0032d40bf60337e63d927b53f
	L2Client, err := wtCommon.GetConnection(simplifiedConfig.L2WebsocketURL, config.Retries)
	if err != nil {
		panic(err)
	}

	blockNumberInt, err := L2Client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	blockNumber := new(big.Int).SetUint64(blockNumberInt) 

	proofOfDilegence := []byte{0x1}
	hash := crypto.Keccak256Hash(proofOfDilegence)
	signedProofOfDiligence, err := vault.SignData(hash.Bytes(), apitypes.TextPlain.Mime)
	if err != nil{
		panic(err)
	}

	wtCommon.Info("watchtower Address: " + auth.From.Hex())
	wtCommon.Info("signature: " + hexutil.Encode(signedProofOfDiligence) )


	var wg sync.WaitGroup
	wg.Add(1)
	SubmitProofToInclusionProofManager(&wg, auth, simplifiedConfig, blockNumber, proofOfDilegence, signedProofOfDiligence, proofSubmissionClient)
	wg.Wait()
}
