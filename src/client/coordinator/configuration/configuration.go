package configuration

import "crypto/ecdsa"

type CoordinatorConfiguration struct {
	PrivateKey      *ecdsa.PrivateKey
	PublicKeyHex    string
	WatchingChainID string
}
