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
	Account      accounts.Account
	backend      accounts.Backend
	transactOpts bind.TransactOpts
}

type VaultConfig struct {
	ChainID      *big.Int
	Address      common.Address
	PrivateKey   *ecdsa.PrivateKey
	Endpoint     string
	EncryptedKey string
	KeyType      string
}

func SetupVault(vc *VaultConfig) (*Vault, error) {
	watchtoweUrl := accounts.URL{}

	watchtowerAccount := accounts.Account{Address: vc.Address, URL: watchtoweUrl}

	if vc.Endpoint != "" {
		endpointBytes, _ := json.Marshal(vc.Endpoint)
		watchtoweUrl.UnmarshalJSON(endpointBytes)
		watchtowerAccount.URL = watchtoweUrl

		backend, err := web3signer.NewExternalBackend(vc.Endpoint)

		if err != nil {
			wtCommon.Info(err)
		} else {
			if vc.Address.Cmp(common.HexToAddress("0")) == 0 {
				if len(backend.Wallets()) == 0 {
					return nil, errors.New("web3signer: no wallet found")
				}

				if len(backend.Wallets()[0].Accounts()) == 0 {
					return nil, errors.New("web3signer: no keys found")
				}

				watchtowerAddress := backend.Wallets()[0].Accounts()[0].Address

				watchtowerAccount = accounts.Account{Address: watchtowerAddress}
			}
			wtCommon.Info("keystore: web3signer: " + watchtowerAccount.URL.String())
			watchtowerAccount.URL = accounts.URL{Scheme: "extapi", Path: vc.Endpoint}
			return &Vault{name: "web3signer", Account: watchtowerAccount, backend: backend}, nil
		}
	}

	if vc.PrivateKey != nil {
		keyType := "raw"
		if len(vc.KeyType) != 0 {
			keyType = vc.KeyType
		}
		backend := newRawBackend(vc.PrivateKey)
		transactOpts := NewRawTransactionOpts(vc.PrivateKey, vc.ChainID)
		watchtowerAccount.URL = accounts.URL{Scheme: keyType, Path: vc.Address.Hex()}
		wtCommon.Info("keystore: " + watchtowerAccount.URL.String())
		return &Vault{name: keyType, Account: watchtowerAccount, backend: backend, transactOpts: *transactOpts}, nil
	} else if len(vc.EncryptedKey) != 0 {
		// corner-case where we could not retrieve encrypted key in config.go initially
		privateKey, err := opCommon.LoadPrivateKey(vc.EncryptedKey, vc.KeyType)
		if err != nil {
			return nil, err
		}
		watchtowerAccount.Address = crypto.PubkeyToAddress(privateKey.PublicKey)
		watchtowerAccount.URL = accounts.URL{Scheme: vc.KeyType, Path: vc.EncryptedKey}
		backend := newRawBackend(privateKey)
		transactOpts := NewRawTransactionOpts(privateKey, vc.ChainID)
		wtCommon.Info("keystore : " + vc.KeyType + " : " + vc.EncryptedKey)
		return &Vault{name: vc.KeyType, Account: watchtowerAccount, backend: backend, transactOpts: *transactOpts}, nil
	}

	wtCommon.Fatal("SetupSigner Failed, please configure watchtower private keys in plaintext, web3signer, or encrypted file system")
	return nil, nil
}

func (vault *Vault) NewTransactOpts(chainID *big.Int) *bind.TransactOpts {
	if vault.name == "raw" || vault.name == "gocryptfs" || vault.name == "w3secretkeys" {
		return &vault.transactOpts
	}
	if vault.name == "web3signer" {
		wallets := vault.backend.Wallets()
		for _, wallet := range wallets {
			if wallet.Contains(vault.Account) {
				return NewWeb3SignerTransactionOpts(wallet, vault.Account, chainID)
			}
		}
	}
	panic("NewTransactOpts()")
	return nil
}

func (vault *Vault) SignData(data []byte, mimeType string) ([]byte, error) {
	wallets := vault.backend.Wallets()

	// there can be more than one wallet, say two usb hardware wallet plugged into a system
	for _, wallet := range wallets {
		if wallet.Contains(vault.Account) {
			signedData, err := wallet.SignData(vault.Account, mimeType, data)
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
		if wallet.Contains(vault.Account) {
			signedTx, err := wallet.SignTx(vault.Account, tx, chainID)
			if err != nil {
				return nil, err
			}
			return signedTx, nil
		}
	}

	wtCommon.Fatal("SignData failed, watchtower account not found in the keystore")
	return nil, nil
}

func GetVaultConfig(config *wtCommon.SimplifiedConfig) *VaultConfig {
	return &VaultConfig{
		Address:      config.WatchtowerAddress,
		PrivateKey:   config.PrivateKey,
		ChainID:      big.NewInt(config.ProofSubmissionChainID),
		Endpoint:     config.ExternalSignerEndpoint,
		EncryptedKey: config.EncryptedKey,
		KeyType:      config.KeyType,
	}
}
