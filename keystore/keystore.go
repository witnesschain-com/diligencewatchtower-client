package keystore

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/witnesschain-com/diligencewatchtower-client/third_party/web3signer"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

type Vault struct {
	name         string
	account      accounts.Account
	backend      accounts.Backend
	transactOpts bind.TransactOpts
}

func SetupVault(watchtowerAddress common.Address, chainID *big.Int, privateKey *ecdsa.PrivateKey, endpoint string) (*Vault, error) {
	watchtoweUrl := accounts.URL{}
	

	watchtowerAccount := accounts.Account{Address: watchtowerAddress, URL: watchtoweUrl}

	if privateKey != nil {
		backend := newRawBackend(privateKey)
		transactOpts := NewRawTransactionOpts(privateKey, chainID)
		watchtowerAccount.URL = accounts.URL{Scheme: "raw", Path: watchtowerAddress.Hex()}
		wtCommon.Info("keystore: " + watchtowerAccount.URL.String())
		return &Vault{name: "raw", account: watchtowerAccount, backend: backend, transactOpts: *transactOpts}, nil
	}

	if endpoint != "" {
		endpointBytes, _ := json.Marshal(endpoint)
		watchtoweUrl.UnmarshalJSON(endpointBytes)
		watchtowerAccount.URL = watchtoweUrl

		backend, err := web3signer.NewExternalBackend(endpoint)

		if err != nil {
			wtCommon.Info(err)
		} else {
			if watchtowerAddress.Cmp(common.HexToAddress("0")) == 0 {
				if len(backend.Wallets()) == 0 {
					return nil, errors.New("web3signer: no wallet found")
				}

				if len(backend.Wallets()[0].Accounts()) == 0 {
					return nil, errors.New("web3signer: no keys found")
				}

				watchtowerAddress = backend.Wallets()[0].Accounts()[0].Address
				
				
				watchtoweUrl := accounts.URL{}
				watchtoweUrl.UnmarshalJSON([]byte(endpoint))
				watchtowerAccount = accounts.Account{Address: watchtowerAddress, URL: watchtoweUrl}
			}
			wtCommon.Info("keystore: web3signer: " + watchtowerAccount.URL.String())
			return &Vault{name: "web3signer", account: watchtowerAccount, backend: backend}, nil
		}
	}

	wtCommon.Fatal("SetupSigner Failed, please configure watchtower private keys in plaintext, web3signer, or encrypted file system")
	return nil, nil
}

func (vault *Vault) NewTransactOpts(chainID *big.Int) *bind.TransactOpts {
	if vault.name == "raw" {
		return &vault.transactOpts
	}
	if vault.name == "web3signer" {
		wallets := vault.backend.Wallets()
		for _, wallet := range wallets{
			if wallet.Contains(vault.account){
				return NewWeb3SignerTransactionOpts(wallet, vault.account, chainID)
			}
		}
	}
	panic("NewTransactOpts()")
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

	wtCommon.Fatal("SignData failed, watchtower account not found in the keystore")
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

	wtCommon.Fatal("SignData failed, watchtower account not found in the keystore")
	return nil, nil
}
