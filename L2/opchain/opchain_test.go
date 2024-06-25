package opchain

import (
	"context"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

func TestParseOutputProposed(t *testing.T) {
	//fetch log from contract

	config := wtCommon.LoadConfigFromJson()
	config.CurrentlyWatchingL2 = "optimism"
	simpleConfig := wtCommon.LoadSimplifiedConfig(config, nil)

	client, err := wtCommon.GetConnection(simpleConfig.L1WebsocketURL, config.Retries)
	if err != nil {
		t.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	// https://sepolia.etherscan.io/tx/0x2c22200a8b55d89cdb780c5fa4ec648ec4421be0032d40bf60337e63d927b53f
	blockNumber = 5321111

	// form filter query for L2OO contract
	query := ethereum.FilterQuery{
		Addresses: []common.Address{simpleConfig.StateCommitmentAddress},
		FromBlock: new(big.Int).SetUint64(blockNumber - 1),
		ToBlock:   new(big.Int).SetUint64(blockNumber),
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		t.Fatal(err)
	}
	wtCommon.Info(logs)

	if len(logs) != 1 {
		t.Fatalf("More than one log. len(logs) = %v", len(logs))
	}

	proposedOutput := ParseOutputProposed(logs[0])
	expectedProposedOutput := "0x39258197d9de61834c671dca234996b177b00b5585658e7acbda94af72f2ca7c"

	if proposedOutput.L2BlockNumber.String() != "8275800" {
		t.Fatalf("%v != 8275800", proposedOutput.L2BlockNumber.String())
	}

	if proposedOutput.OutputRoot != expectedProposedOutput {
		t.Fatalf("proposedOutput %v != %v", proposedOutput.OutputRoot, expectedProposedOutput)
	}

}

func TestGetOutputRootByBlockNumber(t *testing.T) {
	config := wtCommon.LoadConfigFromJson()
	config.CurrentlyWatchingL2 = "optimism"
	simpleConfig := wtCommon.LoadSimplifiedConfig(config, nil)
	client, err := wtCommon.GetConnection(simpleConfig.L2ExecRPCURL, config.Retries)
	if err != nil {
		t.Fatal(err)
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	url := simpleConfig.L2NodeRPCURL
	wtCommon.Info(url)
	hexBlockNumber := strconv.FormatUint(blockNumber, 16)
	root := GetOutputRootByBlockNumber("0x"+hexBlockNumber, url)
	if len(root) == 0 {
		t.Fatal("root is empty.")
	}
	wtCommon.Info(root)

}

func TestTraceBlock(t *testing.T) {
	config := wtCommon.LoadConfigFromJson()
	config.CurrentlyWatchingL2 = "optimism"
	simpleConfig := wtCommon.LoadSimplifiedConfig(config, nil)

	client, err := wtCommon.GetConnection(simpleConfig.L2ExecRPCURL, config.Retries)
	if err != nil {
		t.Fatal(err)
	}

	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}

	roots := TraceBlock(latestHeader.Hash().String(), simpleConfig)
	if len(roots) <= 0 {
		t.Fatalf("len(roots) = %v", len(roots))
	}

	if len(roots[0]) != 66 {
		t.Fatalf("len(roots[0]) = %v", len(roots))
	}
}
