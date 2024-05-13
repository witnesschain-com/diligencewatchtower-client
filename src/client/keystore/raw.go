package keystore

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

var notImplemented = errors.New("Not implemented. Operation not supported.")

type rawBackend struct{
	wallet []accounts.Wallet
}

func (backend *rawBackend) Wallets() []accounts.Wallet {
	return backend.wallet
}

func (backend *rawBackend) Subscribe(sink chan<- accounts.WalletEvent) event.Subscription {
	return event.NewSubscription(func(quit <-chan struct{}) error {
		<-quit
		return nil
	})
}

// raw Wallet
type raw struct {
	key *ecdsa.PrivateKey
	address common.Address
}

func newRawBackend(privateKey *ecdsa.PrivateKey) *rawBackend{
	address, err := wtCommon.GetPublicKeyAddressFromPrivateKey(privateKey)
	if err != nil {
		wtCommon.Fatal(err)
	}
	wallet := &raw{key: privateKey, address: address}
	
	return &rawBackend{wallet: []accounts.Wallet{wallet}}
}

func NewRawTransactionOpts(privateKey *ecdsa.PrivateKey, chainID  *big.Int) *bind.TransactOpts{
	transactionOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		wtCommon.Error(err)
	}
	return transactionOpts
}


func (api *raw) SignData (account accounts.Account, mimeType string, data []byte) ([]byte, error){
	ethereumMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	hash := crypto.Keccak256Hash([]byte(ethereumMessage))
	signature, err := crypto.Sign(hash.Bytes(), api.key)
	if err != nil {
		return nil, err
	}
	signature[64] += 27
	return signature, nil
}

func (api *raw) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {


	auth := types.LatestSignerForChainID(chainID)
	
	signedTx , err := types.SignTx(tx, auth, api.key)
	if err != nil {
		wtCommon.Fatal(err)
	}
	return signedTx, nil
}

func (api *raw) Accounts() []accounts.Account {
	wtCommon.Fatal(notImplemented)
	return []accounts.Account{}

}

func (api *raw) URL() accounts.URL{
	wtCommon.Fatal(notImplemented)
	return accounts.URL{}
}

func (api *raw) Status() (string, error){
	wtCommon.Fatal(notImplemented)
	return "O.K.", nil
}

func (api *raw) Open(passphrase string) error{
	wtCommon.Fatal(notImplemented)
	return nil
}

func (api *raw) Close() error{
	wtCommon.Fatal(notImplemented)
	return nil
}

func (api *raw) Contains(account accounts.Account) bool {
	wtCommon.Fatal(notImplemented)
	if api.address == account.Address{
		return true
	}
	return false
}

func (api *raw) Derive(path accounts.DerivationPath, pin bool) (accounts.Account, error){
	wtCommon.Fatal(notImplemented)
	return accounts.Account{}, nil
}

func (api *raw) SelfDerive(bases []accounts.DerivationPath, chain ethereum.ChainStateReader){
	wtCommon.Fatal(notImplemented)
}

func (api *raw) SignDataWithPassphrase(account accounts.Account, passphrase, mimeType string, data []byte) ([]byte, error) {
	wtCommon.Fatal(notImplemented)
	return []byte{0}, nil
}

func (api *raw) SignText(account accounts.Account, text []byte) ([]byte, error) {
	wtCommon.Fatal(notImplemented)
	return []byte{}, nil
}

func (api *raw) SignTextWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error){
	wtCommon.Fatal(notImplemented)
	return []byte{}, nil
}

func (api *raw) SignTxWithPassphrase(account accounts.Account, passphrase string, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error){
	wtCommon.Fatal(notImplemented)
	return &types.Transaction{}, nil
}

