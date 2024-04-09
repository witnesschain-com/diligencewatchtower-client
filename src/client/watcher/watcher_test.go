package watcher

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"

	wtCommon "github.com/diligencewatchtower-client/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestProcessIntermediateBlocks(t *testing.T) {
	config := wtCommon.LoadConfigFromJson()
	config.CurrentlyWatchingL2 = "optimism"
	simpleConfig := wtCommon.LoadSimplifiedConfig(config, nil)
	// L1 trx https://sepolia.etherscan.io/tx/0x2c22200a8b55d89cdb780c5fa4ec648ec4421be0032d40bf60337e63d927b53f
	// L2 block https://sepolia-optimism.etherscan.io/block/8275800
	L2BlockNumber, _ := new(big.Int).SetString("8275800", 10)
	ultimateBlockHash, penultimateBlockHash, L2OutputRoot := ProcessIntermediateBlocks(L2BlockNumber, simpleConfig)

	expectedUltimateBlockHash := "0x103fe7c41867dac7e9ff13b3497d92a2302d593873fa76dfd65f07f4518cbdca"
	expectedPenultimateBlockHash := "0xf70355e398e4cfe38d2907a16b566a9b1f857dd4d206dca1c859ce73b7dae244"
	expectedL2OutputRoot := "0x39258197d9de61834c671dca234996b177b00b5585658e7acbda94af72f2ca7c"

	if ultimateBlockHash != expectedUltimateBlockHash {
		t.Fatalf("ultimateBlockHash %v!=%v", ultimateBlockHash, expectedUltimateBlockHash)
	}

	if penultimateBlockHash != expectedPenultimateBlockHash {
		t.Fatalf("penultimateBlockHash %v!=%v", penultimateBlockHash, expectedPenultimateBlockHash)
	}

	if L2OutputRoot != expectedL2OutputRoot {
		t.Fatalf("L2OutputRoot %v!=%v", L2OutputRoot, expectedL2OutputRoot)
	}
}

func TestFetchIntermediateStateRoots(t *testing.T) {

	config := wtCommon.LoadConfigFromJson()
	config.CurrentlyWatchingL2 = "optimism"
	simpleConfig := wtCommon.LoadSimplifiedConfig(config, nil)
	ultimateBlockHash := "0x103fe7c41867dac7e9ff13b3497d92a2302d593873fa76dfd65f07f4518cbdca"
	penultimateBlockHash := "0xf70355e398e4cfe38d2907a16b566a9b1f857dd4d206dca1c859ce73b7dae244"
	midPointPenultimateBlock, midPointUltimate := FetchIntermediateStateRoots(
		simpleConfig,
		ultimateBlockHash,
		penultimateBlockHash)

	expectedMidPointPenultimateBlock := "0x7bd00f11c18429ed41add9266884ace0e7e0334e37abbd00d7a037db77d4355c"
	expectedMidPointUltimate := "0x41641c9c7b67910c2856db1aa7f83a6131793d4ade96c703ac0ad5fad24d4bb2"

	if midPointPenultimateBlock != expectedMidPointPenultimateBlock {
		t.Fatalf("midPointPenultimateBlock %v!=%v", midPointPenultimateBlock, expectedMidPointPenultimateBlock)
	}

	if midPointUltimate != expectedMidPointUltimate {
		t.Fatalf("midPointUltimate %v!=%v", midPointUltimate, expectedMidPointPenultimateBlock)
	}
}

func TestComputeProofOfDiligence(t *testing.T) {
	//L1 submit proof https://sepolia.etherscan.io/tx/0x430b4f5bd295520c895ab1923829ad20b5a3cc9b2d3ee496af28618bb0bafff5
	l2latestBlockNumber := new(big.Int).SetUint64(8275800)
	midPointPenultimateBlock := "0x7bd00f11c18429ed41add9266884ace0e7e0334e37abbd00d7a037db77d4355c"
	midPointUltimate := "0x41641c9c7b67910c2856db1aa7f83a6131793d4ade96c703ac0ad5fad24d4bb2"
	proofOfDilegence := ComputeProofOfDiligence(l2latestBlockNumber, midPointPenultimateBlock, midPointUltimate)
	expectedProofOfDilegence, err := hex.DecodeString("383237353830302c3078376264303066313163313834323965643431616464393236363838346163653065376530333334653337616262643030643761303337646237376434333535632c3078343136343163396337623637393130633238353664623161613766383361363133313739336434616465393663373033616330616435666164323464346262322c30")

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(proofOfDilegence, expectedProofOfDilegence) != true {
		t.Fatalf("proofOfDilegence %v!=%v", proofOfDilegence, expectedProofOfDilegence)
	}
}

func TestSignProofOfDiligence(t *testing.T) {
	proofOfDilegence, err := hex.DecodeString("383237353830302c3078376264303066313163313834323965643431616464393236363838346163653065376530333334653337616262643030643761303337646237376434333535632c3078343136343163396337623637393130633238353664623161613766383361363133313739336434616465393663373033616330616435666164323464346262322c30")
	if err != nil {
		t.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("ebe3db0a12a2b2f249ea8c699aed3bf542a5a22d6d42fb98c4f715d244ccb5fe")
	if err != nil {
		t.Fatal(err)
	}

	signedProofOfDiligence := SignProofOfDiligence(proofOfDilegence, privateKey)

	expectedSignedProofOfDiligence, err := hex.DecodeString("0059d9cf0369f80525bf01d7efa8833b090a9729176c8912a6b167300e8fb82d54d7edea0d3b64c68d1c0c48cab09b9621ede685deb2bff2d203130301a3e4701b")
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(signedProofOfDiligence, expectedSignedProofOfDiligence) != true {
		t.Fatalf("signedProofOfDiligence %v!=%v", signedProofOfDiligence, expectedSignedProofOfDiligence)
	}
}
