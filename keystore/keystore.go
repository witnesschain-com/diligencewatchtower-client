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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/witnesschain-com/diligencewatchtower-client/third_party/web3signer"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	opCommon "github.com/witnesschain-com/operator-cli/common"
)

type Vault struct {
	name         string
	account      accounts.Account
	backend      accounts.Backend
	transactOpts bind.TransactOpts
}

type VaultConfig struct {
	ChainID *big.Int
	WatchtowerAddress common.Address
	PrivateKey *ecdsa.PrivateKey
	Endpoint string
	GocryptfsKey string
}

func SetupVault(watchtowerAddress common.Address, chainID *big.Int, privateKey *ecdsa.PrivateKey, endpoint string) (*Vault, error) {
	gocryptfsKey := "/home/kripashanker/.witnesschain.com/.encrypted_keys/wt1"
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
				
				
				watchtowerAccount = accounts.Account{Address: watchtowerAddress}
			}
			wtCommon.Info("keystore: web3signer: " + watchtowerAccount.URL.String())
			watchtowerAccount.URL = accounts.URL{Scheme: "extapi", Path: endpoint}
			return &Vault{name: "web3signer", account: watchtowerAccount, backend: backend}, nil
		}
	}

	if len(gocryptfsKey) != 0 {
		privateKey, err := opCommon.LoadPrivateKey(gocryptfsKey)
		if err != nil {
			return nil, err
		}
		watchtowerAccount.Address = crypto.PubkeyToAddress(privateKey.PublicKey)
		watchtowerAccount.URL = accounts.URL{Scheme: "gocryptfs", Path: gocryptfsKey}
		backend := newRawBackend(privateKey)
		wtCommon.Info("keystore: gocryptfs: " + gocryptfsKey)
		return &Vault{name: "gocryptfs", account: watchtowerAccount, backend: backend}, nil
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

	// there can be more than one wallet, say two usb hardware wallet plugged into a system
	for _, wallet := range wallets {
		if wallet.Contains(vault.account) {
			signedData, err := wallet.SignData(vault.account, "plain/text", data)
			if err != nil {
				wtCommon.Error(err)
				return nil, err
			}
			return signedData, nil
		}
	}

	wtCommon.Fatal("SignData failed, watchtower account not found in any wallet.")

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
