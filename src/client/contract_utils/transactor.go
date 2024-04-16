package contractutils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

func SetGasPrice(
	config *wtCommon.SimplifiedConfig,
) *big.Int {
	wtCommon.Info("Setting `gasPrice` on L1 ...")

	// gasPrice = nil => eistmate gas price
	var gasPrice *big.Int = nil
	if config.GasPrice != 0 {
		gasPrice = big.NewInt(config.GasPrice)
		wtCommon.Success(gasPrice)
	} else {
		wtCommon.Success("`gasPrice` set to estimate from the network")
	}

	return gasPrice
}

func GetNonce(
	L1Client *ethclient.Client,
	publicKeyAddress common.Address,
) (uint64, error) {
	wtCommon.Info("Getting `nonce` on L1 ...")

	nonce, err := L1Client.PendingNonceAt(context.Background(), publicKeyAddress)
	if err != nil {
		wtCommon.Fatal(err)
	}

	if nonce == 0 {
		wtCommon.Warning("Got nonce as `0` on L1!")
	}

	wtCommon.Success(nonce)
	return nonce, nil
}

func GetTransactor(
	privateKey *ecdsa.PrivateKey,
	chainID int64,
	nonce uint64,
	gasPrice *big.Int,
) (*bind.TransactOpts, error) {

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)        // in wei
	auth.GasLimit = uint64(3_000_000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

func GetTransactOpts(
	config *wtCommon.SimplifiedConfig,
	SubmissionChainClient *ethclient.Client,

) (*bind.TransactOpts, error) {

	// get eth address associated with watchtower private key stored in the config
	publicKeyAddress, err := wtCommon.GetPublicKeyAddressFromPrivateKey(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	// set gas price for submittign PoD
	gasPrice := SetGasPrice(config)

	// fetch nonce associated with this watchtower eth address on proof submission chain
	nonce, err := GetNonce(SubmissionChainClient, publicKeyAddress)
	if err != nil {
		return nil, err
	}

	// prepare transactor object to invoke the transaction as watchtower
	auth, err := GetTransactor(
		config.PrivateKey,
		config.ProofSubmissionChainID,
		nonce,
		gasPrice,
	)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
