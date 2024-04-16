package watcher

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/cbergoon/merkletree"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
)

type InclusionProof struct {
	Receipt   opchain.TransactionReceipt // receipt
	Signature []byte                     // signed receipt
}

// InclusionProofHash implements the Content interface provided by merkletree
type InclusionProofHash struct {
	proofHash string
}

// CalculateHash hashes the values of a InclusionProofHash
func (inclusionProofHash InclusionProofHash) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(inclusionProofHash.proofHash)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (inclusionProofHash InclusionProofHash) Equals(other merkletree.Content) (bool, error) {
	otherIPFHash, err := other.(InclusionProofHash)
	if !err {
		return false, errors.New("value is not of type InclusionProofHash")
	}
	return inclusionProofHash.proofHash == otherIPFHash.proofHash, nil
}

// InclusionProofBatch
type InclusionProofBatch struct {
	batchStart    int64
	batchSize     int64
	currentProofs []merkletree.Content
}

// append a new proof to the batch
func (inclusionProofBatch *InclusionProofBatch) append(newProofHash InclusionProofHash) {
	if len(inclusionProofBatch.currentProofs) < int(inclusionProofBatch.batchSize) {
		inclusionProofBatch.currentProofs = append(inclusionProofBatch.currentProofs, newProofHash)
		wtCommon.Info("inclusion proof appended to the batch ...")
	} else {
		wtCommon.Error("appending to a full batch!")
	}
}

// reset (empty) the current batch
func (inclusionProofBatch *InclusionProofBatch) resetBatch() {
	inclusionProofBatch.currentProofs = nil
	inclusionProofBatch.batchStart = -1
	wtCommon.Info("Batch reset!")
}

// check if the batch is full
func (inclusionProofBatch *InclusionProofBatch) checkBatchCompletion() bool {
	return len(inclusionProofBatch.currentProofs) == int(inclusionProofBatch.batchSize)
}

// returns the root hash of the merkle tree formed by the inclusion proofs
func (inclusionProofBatch *InclusionProofBatch) getRootHash() []byte {
	inclusionProofTree, err := merkletree.NewTree(inclusionProofBatch.currentProofs)
	if err != nil {
		wtCommon.Error(err)
	}
	merkleRoot := inclusionProofTree.MerkleRoot()
	wtCommon.Info("Batch merkle root is " + hex.EncodeToString(merkleRoot))
	return merkleRoot
}
