package watcher

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
	"github.com/witnesschain-com/diligencewatchtower-client/keystore"
)

func FetchIntermediateStateRoots(
	config *wtCommon.SimplifiedConfig,
	ultimateBlockHash string,
	penultimateBlockHash string,

) (string, string) {
	wtCommon.Info("Running trace on L2 RPC node ...")

	traceResultUltimate := opchain.TraceBlock(ultimateBlockHash, config)
	traceResultPenultimate := opchain.TraceBlock(penultimateBlockHash, config)

	wtCommon.Success("Found `" + strconv.Itoa(len(traceResultPenultimate)) + " and " + strconv.Itoa(len(traceResultUltimate)) + "` intermediate state roots")

	midPointPenultimateBlock := traceResultPenultimate[len(traceResultPenultimate)/2]
	midPointUltimate := traceResultUltimate[len(traceResultUltimate)/2]

	wtCommon.Info("`midPoint` of stateRoots on L2 are ...")
	wtCommon.Success(midPointPenultimateBlock)
	wtCommon.Success(midPointUltimate)
	return midPointPenultimateBlock, midPointUltimate
}

func ComputeProofOfDiligence(
	latestBlockNumber *big.Int,
	midPointPenultimateBlock string,
	midPointUltimateBlock string,

) []byte {

	version_number := big.NewInt(0)
	wtCommon.Info("the `version_number` for proof of diligence is")
	wtCommon.Success(version_number)
	wtCommon.Info("Computing proof of diligence")

	proofOfDilegence := []byte(latestBlockNumber.String() + "," + midPointPenultimateBlock + "," + midPointUltimateBlock + "," + version_number.String())

	return proofOfDilegence
}

func SignProofOfDiligence(
	proofOfDilegence []byte,
	watchtower common.Address,
	vault *keystore.Vault,
) []byte {
	hash := crypto.Keccak256Hash(proofOfDilegence)

	prefix := []byte("\x19Ethereum Signed Message:\n32")
	var hash_bytes []byte = hash.Bytes()

	proofOfDilegenceToHash := append(prefix, hash_bytes...)
	final_hash := crypto.Keccak256Hash(proofOfDilegenceToHash)

	signatureOfProofOfDiligence, err := vault.SignData(final_hash[:])
	if err != nil {
		wtCommon.Error(err)
	}

	wtCommon.Success("Successfully signed Proof of Diligence")
	return signatureOfProofOfDiligence
}
