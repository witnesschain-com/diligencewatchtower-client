package keystore

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Following code is taken from clef of go-ethereum project

// NewWeb3signerTransactor is a utility method to easily create a transaction signer
// with a web3signer backend.

func NewWeb3SignerTransactionOpts(wallet accounts.Wallet, account accounts.Account, chainID *big.Int) *bind.TransactOpts {
	
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, transaction, chainID) 
		},
		Context: context.Background(),
	}
}
