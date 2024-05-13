package keystore

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/witnesschain-com/diligencewatchtower-client/watchtower/third_party/web3signer"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

type Vault struct {
	name string
	account accounts.Account
	backend accounts.Backend
	transactOpts bind.TransactOpts
}

func SetupVault(config *wtCommon.SimplifiedConfig) (*Vault, error) {

	account := accounts.Account{Address: config.WatchtowerAddress}

	if config.ExternalSignerEndpoint != "" {
		backend, err := web3signer.NewExternalBackend(config.ExternalSignerEndpoint)
		if err != nil {
			wtCommon.Error(err)
		}
		return &Vault{name: "web3signer", account:  account, backend: backend}, nil
	}

	if config.Vault != "" {
		wtCommon.Info("setup encrypted vault")
	}

	if config.PrivateKey!= nil {
		backend := newRawBackend(config.PrivateKey)
		chainID := new(big.Int).SetUint64(uint64(config.ProofSubmissionChainID))
		transactOpts := NewRawTransactionOpts(config.PrivateKey, chainID)
		return &Vault{name: "raw", account: account, backend: backend, transactOpts: *transactOpts}, nil
	}

	wtCommon.Fatal("SetupSigner Failed");

	return nil, nil
}

func (vault *Vault) NewTransactOpts(chainID *big.Int) *bind.TransactOpts{
	if vault.name == "raw"{
		return &vault.transactOpts
	}
	if vault.name == "web3signer"{
		return NewWeb3SignerTransactionOpts(vault, nil)
	}
	panic("NewTransactOpts()");
	return nil
}


func (vault *Vault) SignData(data []byte) ([]byte, error) {
	wallets := vault.backend.Wallets()

	for _, wallet := range wallets {
		signedData, err := wallet.SignData(vault.account, "plain/text", data)
		if err != nil {
			wtCommon.Error(err)
			return nil, err
		}
		return signedData, nil
	}

	wtCommon.Fatal("SignData failed, watchtower account not found in the keystore");
	return nil, nil
}

func (vault *Vault) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {

	wallets := vault.backend.Wallets()

	for _, wallet := range wallets {
		signedTx, err := wallet.SignTx(vault.account, tx, chainID)
		if err != nil {
			return nil, err
		}
		return signedTx, nil
	}
	
	wtCommon.Fatal("SignData failed, watchtower account not found in the keystore");
	return nil, nil
}

