/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package opchain

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/core/types"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

// Parses output proposed event log on L2OO
func ParseOutputProposed(logEvent types.Log) *OutputProposed {

	wtCommon.Info("Processing proposal at L2 Output Oracle")
	proposedOutput := new(OutputProposed)
	topics := logEvent.Topics

	wtCommon.Info("Getting `outputRoot` from L2 Output Oracle ...")
	proposedOutput.OutputRoot = topics[1].Hex()
	wtCommon.Success(proposedOutput.OutputRoot)

	wtCommon.Info("Getting `L2OutputIndex` from L2 Output Oracle ...")
	proposedOutput.L2OutputIndex = topics[2].Hex()
	wtCommon.Success(proposedOutput.L2OutputIndex)

	wtCommon.Info("Getting `L2BlockNumber` from L2 Output Oracle ...")
	proposedOutput.L2BlockNumber = topics[3].Big()
	wtCommon.Success(proposedOutput.L2BlockNumber)

	wtCommon.Info("Getting `L1Timestamp` from L2 Output Oracle ...")
	proposedOutput.L1Timestamp = hex.EncodeToString(logEvent.Data)
	wtCommon.Success(proposedOutput.L1Timestamp)

	return proposedOutput
}
