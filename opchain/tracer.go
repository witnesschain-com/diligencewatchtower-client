/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package opchain

import (
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

// Requests the node to re-run the block
// generates intermediate state roots after each transaction in the block
// params: blockhash of the block to re-run and trace
// return: a list of state root after each transaction in the block
func TraceBlock(blockHash string, config *wtCommon.SimplifiedConfig) []string {

	// Connect to the RPC node via http
	client, err := wtCommon.GetRPCConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer client.Close()

	// interface for result of the trace roots
	// `result` will hold the list of intermediate state roots from the block
	var result []string

	// make the tracer call and collect result
	err = client.Call(&result, "debug_intermediateRoots", blockHash)
	if err != nil {
		// TODO: change to FATAL to ERROR, return error and handle it
		wtCommon.Fatal(err)
	}

	return result
}
