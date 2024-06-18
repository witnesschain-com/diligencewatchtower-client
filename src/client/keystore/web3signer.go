package keystore

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

// Following code is taken from clef of go-ethereum project

// NewWeb3signerTransactor is a utility method to easily create a transaction signer
// with a web3signer backend.

func NewWeb3SignerTransactionOpts(vault *Vault, chainID *big.Int) *bind.TransactOpts {
	wallets := vault.backend.Wallets()
	if len(wallets) != 1 {
		wtCommon.Error("Web3signer: more than one wallet found")
	}

	return &bind.TransactOpts{
		From: vault.account.Address,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			if address != vault.account.Address {
				return nil, bind.ErrNotAuthorized
			}
			return wallets[0].SignTx(vault.account, transaction, nil) // web3signer enforces its own chain id
		},
		Context: context.Background(),
	}
}
