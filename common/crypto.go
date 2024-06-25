/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package common

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetPublicKeyAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		Fatal("Could not get publicKeyECDSA")
		return common.BigToAddress(big.NewInt(0)), errors.New("Could not get publicKeyECDSA")
	}

	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return publicKeyAddress, nil
}

func RecommendKey() {
	randomPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		Error(err)
	}

	randomPrivateKeyBytes := crypto.FromECDSA(randomPrivateKey)
	randomPrivateKeyHex := hexutil.Encode(randomPrivateKeyBytes)[2:]

	Info("For example: you may consider using the below random `private_key` ...")
	Success(randomPrivateKeyHex)
}
