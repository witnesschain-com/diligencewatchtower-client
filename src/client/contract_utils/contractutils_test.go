package contractutils

// "encoding/hex"
// "math/big"
// "testing"

// wtCommon "github.com/diligencewatchtower-client/common"

// Can only submit proof for the latest block number
// func TestSubmitProofToDiligenceProofManager(t *testing.T){

// 	config := wtCommon.LoadConfigFromJson()
// 	config.CurrentlyWatchingL2 = "optimism"
// 	simplifiedConfig := wtCommon.LoadSimplifiedConfig(config, nil)
// 	simplifiedConfig.L1WebsocketURL = "http://localhost:8545"
// 	wtCommon.Info(simplifiedConfig)

// 	publicKeyAddress, err := wtCommon.GetPublicKeyAddressFromPrivateKey(config.PrivateKey)
// 	if err != nil {
// 		wtCommon.Fatal(err)
// 	}
// 	L1Client, err := wtCommon.GetConnection(simplifiedConfig.L1WebsocketURL, config.Retries)
// 	if err != nil {
// 		wtCommon.Fatal(err)
// 	}

// 	// set gas price for submittign PoD
// 	gasPrice := SetGasPrice(simplifiedConfig)

// 	// fetch nonce associated with this watchtower eth address on chain
// 	nonce, err := GetNonce(L1Client, publicKeyAddress)
// 	if err != nil {
// 		wtCommon.Fatal(err)
// 	}

// 	// prepare transactor object to invoke the transaction as watchtower
// 	auth, err := GetTransactor(
// 		config.PrivateKey,
// 		simplifiedConfig.CurrentL1ChainID,
// 		nonce,
// 		gasPrice,
// 	)
// 	if err != nil {
// 		wtCommon.Fatal(err)
// 	}

// 	// https://sepolia.etherscan.io/tx/0x2c22200a8b55d89cdb780c5fa4ec648ec4421be0032d40bf60337e63d927b53f
// 	blockNumber, status := new(big.Int).SetString("5321111", 10)
// 	if status == false {
// 		t.Fatal(err)
// 	}

// 	proofOfDilegence, err := hex.DecodeString("383237353830302c3078376264303066313163313834323965643431616464393236363838346163653065376530333334653337616262643030643761303337646237376434333535632c3078343136343163396337623637393130633238353664623161613766383361363133313739336434616465393663373033616330616435666164323464346262322c30")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	signedProofOfDiligence, err := hex.DecodeString("0059d9cf0369f80525bf01d7efa8833b090a9729176c8912a6b167300e8fb82d54d7edea0d3b64c68d1c0c48cab09b9621ede685deb2bff2d203130301a3e4701b")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	transaction, err := SubmitProofToDiligenceProofManager(auth, simplifiedConfig, blockNumber, proofOfDilegence, signedProofOfDiligence, L1Client)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	wtCommon.Info(transaction)
// 	wtCommon.Info(transaction.Hash())
// 	_, err = WaitForTransactionReceipt(
// 		simplifiedConfig,
// 		transaction,
// 		L1Client,
// 	)
// }

// func SubmitProofToDiligenceProofManager(
// 	auth *bind.TransactOpts,
// 	config *wtCommon.SimplifiedConfig,
// 	blockNumber *big.Int,
// 	proofOfDilegence []byte,
// 	signatureOfProofOfDiligence []byte,
// 	L1Client *ethclient.Client,
// ) (*types.Transaction, error) {
