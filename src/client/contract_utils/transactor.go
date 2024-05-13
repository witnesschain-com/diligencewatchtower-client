package contractutils

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/watchtower/keystore"
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
	vault *keystore.Vault,

) (*bind.TransactOpts, error) {


	chainID, err := SubmissionChainClient.ChainID(context.Background())
	if err != nil {
		wtCommon.Fatal(err)
	}

	auth := vault.NewTransactOpts(chainID)

	return auth, nil
}
