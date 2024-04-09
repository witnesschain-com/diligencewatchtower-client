/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package rotation

import (
	"context"

	// "crypto/ecdsa"
	"math/big"
	"sort"
	"strings"

	// "strconv"

	"github.com/diligencewatchtower-client/bindings"
	wtCommon "github.com/diligencewatchtower-client/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Rollup struct {
	ChainID     int64
	Name        string
	WatchTowers int
}

var assignedChain *Rollup = nil
var roundLength int64 = 10

var chianIDToName map[int64]string = map[int64]string{
	993:   "wc",
	84531: "base-goerli",
	420:   "op-goerli",
}

// Returns addresses of registered watchTowers
func getWatchTowers(client *ethclient.Client, config *wtCommon.WatchTowerConfig) ([]common.Address, error) {

	a := common.HexToAddress(config.OperatorRegistry)
	operatorRegistry, err := bindings.NewOperatorRegistry(a, client)
	if err != nil {
		wtCommon.Info("NewOperatorRegistry")
		wtCommon.Error(err)
		return nil, err
	}

	addresses, err := operatorRegistry.GetAllWatchtowerAddresses(nil)
	if err != nil {
		wtCommon.Warning("Unable to retrive addreses")
		wtCommon.Error(err)
		return nil, err
	}
	return addresses, nil
}

func getRollups(client *ethclient.Client, config *wtCommon.WatchTowerConfig) []Rollup {
	a := common.HexToAddress(config.RollupRegistry)
	rollupRegistry, err := bindings.NewRollupRegistry(a, client)
	if err != nil {
		wtCommon.Error("Unable to get instance of NewRollupRegistry")
		wtCommon.Error(err)
	}

	rawRollups, err := rollupRegistry.GetRollups(nil)
	if err != nil {
		wtCommon.Error(err)
	}

	var rollups []Rollup

	for _, rawRollup := range rawRollups {
		chainID := rawRollup[0].Int64()
		name := chianIDToName[chainID]
		count := int(rawRollup[1].Int64())
		if err != nil {
			wtCommon.Error("unable to convert rollup count from string to int")
		}
		rollups = append(rollups, Rollup{chainID, name, count})

	}
	return rollups
}

func findRank(blockhash common.Hash, watchTowerAddresses []common.Address, clientAddress common.Address) int {

	var hashes []string

	// Find hash corresponding to each watchTowers
	for _, a := range watchTowerAddresses {
		// h := hash(blockhash | address)
		d := strings.ToLower(blockhash.Hex()) + "," + strings.ToLower(a.Hex())
		h := strings.ToLower(crypto.Keccak256Hash([]byte(d)).Hex())
		hashes = append(hashes, h)
	}

	sort.Strings(hashes)

	d := strings.ToLower(blockhash.Hex()) + "," + strings.ToLower(clientAddress.Hex())
	myHash := strings.ToLower(crypto.Keccak256Hash([]byte(d)).Hex())

	return sort.SearchStrings(hashes, myHash)
}

func assignChain(hashRank int, rollups []Rollup) *Rollup {
	// assign the selected watchTower to one of the L2 chain
	begin := 0
	assignedChain = nil

	for _, rollup := range rollups {
		rollup.Name = strings.ToLower(rollup.Name)
	}

	sort.Slice(rollups, func(i, j int) bool {
		return rollups[i].Name < rollups[j].Name
	})

	for i, rollup := range rollups {
		end := begin + rollup.WatchTowers
		if hashRank < end {
			assignedChain = &rollups[i]
			break
		}
		begin = end
	}

	return assignedChain
}

/*
* This function selects a set of watchTowrs from registered watchTowers and
* returns chainID of the L2 chain on this watchTower will act as watchTower for
* this round. If it returns -1, then this watchTower is not selected to be watchTower in this round
 */

func Rotate() (*Rollup, error) {

	config := wtCommon.LoadConfigFromJson()

	priv := config.PrivateKey

	client, err := wtCommon.GetConnection(config.EthTestnetWebsocketURL, config.Retries)
	if err != nil {
		wtCommon.Error("Unable to dial")
		wtCommon.Error(err)
		return nil, err
	}
	defer client.Close()

	// latest block added to chain
	latestBlock, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		wtCommon.Error("Unable to get latest block")
		wtCommon.Error(err)
		return nil, err
	}

	currentRound := new(big.Int)
	currentRound.Div(latestBlock.Number(), big.NewInt(roundLength)).Mul(currentRound, big.NewInt(roundLength))

	wtCommon.Info("Got `latestBlockNumber` as")
	wtCommon.Success(latestBlock.Number())

	wtCommon.Info("Got `currentRound` as")
	wtCommon.Success(currentRound)

	if currentRound.Cmp(latestBlock.Number()) != 0 && (assignedChain != nil) {
		return assignedChain, nil
	}

	checkPointBlock, err := client.BlockByNumber(context.Background(), currentRound)

	if err != nil {
		wtCommon.Warning("Unable to get check point block")
		wtCommon.Error(err)
		return nil, err
	}

	addresses, err := getWatchTowers(client, config)
	if err != nil {
		return nil, err
	}

	watchTowerClientAddress := crypto.PubkeyToAddress(priv.PublicKey)
	hashRank := findRank(checkPointBlock.Hash(), addresses, watchTowerClientAddress)

	wtCommon.Info("Hash Rank is")
	wtCommon.Success(hashRank)

	rollups := []Rollup{{420, "optimism", 1}, {84531, "base", 1}}
	_ = rollups

	// TODO: fetch from contracts once it's reqady
	// rollups := getRollups(client, config)

	assignedChain := assignChain(hashRank, rollups)

	return assignedChain, nil
}
