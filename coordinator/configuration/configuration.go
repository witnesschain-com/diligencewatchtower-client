package configuration

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
)

type CoordinatorConfiguration struct {
	WatchtowerAddress      common.Address
	Vault         *keystore.Vault
	WatchingChainID string
}
