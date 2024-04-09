package watcher

import (
	"context"
	"math/big"

	wtCommon "github.com/diligencewatchtower-client/common"
	"github.com/diligencewatchtower-client/opchain"
)

func ProcessIntermediateBlocks(
	L2BlockNumber *big.Int,
	config *wtCommon.SimplifiedConfig,
) (ultimateBlockHash string, penultimateBlockHash string, L2OutputRoot string) {
	wtCommon.Info("Establishing a connection with L2 node")
	L2Client, err := wtCommon.GetConnection(config.L2ExecRPCURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}
	defer L2Client.Close()
	wtCommon.Success("Connected to L2 successfully")

	wtCommon.Info("Getting `blockHash` of `latestBlockNumber` on L2 ...")
	block, err := L2Client.HeaderByNumber(context.Background(), L2BlockNumber)
	if err != nil {
		wtCommon.Fatal(err)
	}

	ultimateBlockHash = block.Hash().Hex()
	wtCommon.Success(ultimateBlockHash)

	wtCommon.Info("Getting `blockHash` of penultimate block on L2 ...")
	penultimateBlockHash = block.ParentHash.Hex()
	wtCommon.Success(penultimateBlockHash)

	wtCommon.Info("Getting output root from L2 node")
	L2OutputRoot = opchain.GetOutputRootByBlockNumber("0x"+L2BlockNumber.Text(16), config.L2NodeRPCURL)
	wtCommon.Success(L2OutputRoot)

	return ultimateBlockHash, penultimateBlockHash, L2OutputRoot

}
