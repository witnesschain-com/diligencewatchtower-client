/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package L1

import (
	"context"
	"time"

	"github.com/diligencewatchtower-client/bindings"
	wtCommon "github.com/diligencewatchtower-client/common"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

func ListenForProposals(config *wtCommon.SimplifiedConfig) (*bindings.L2OutputOracleCaller, ethereum.Subscription, chan types.Log, error) {

	// Connect to the RPC node via web sockets
	client, err := wtCommon.GetConnection(config.L1WebsocketURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}

	L2OOContract, err := bindings.NewL2OutputOracleCaller(config.StateCommitmentAddress, client)
	if err != nil {
		wtCommon.Fatal(err)
	}

	// form filter query for L2OO contract
	query := ethereum.FilterQuery{
		Addresses: []common.Address{config.StateCommitmentAddress},
	}

	// subscribe to the results of streaming filter query
	logs := make(chan types.Log)

	// resubscribe upton error automatically after 2 seconds
	sub := event.Resubscribe(2*time.Second, func(ctx context.Context) (event.Subscription, error) {
		wtCommon.Info("Subscribing to State commitment events on L1")
		eth_sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
		if err == nil {
			wtCommon.Success("Subscribed to the state commitments successfully!")
			wtCommon.Info("Waiting for proposal ...")
		}
		return eth_sub, err
	})

	return L2OOContract, sub, logs, nil

}
