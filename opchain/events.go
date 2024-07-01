/*
	Copyright (c) 2024 Kaleidoscope Blockchain Inc.

	Unless specified otherwise, this work is licensed under the Creative Commons
	Attribution-NonCommercial 4.0 International License.

	To view a copy of this license, visit:
		http://creativecommons.org/licenses/by-nc/4.0/
*/

package opchain

import (
	"context"
	"time"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

func ListenForHead(config *wtCommon.SimplifiedConfig) (ethereum.Subscription, chan *types.Header, error) {

	// Connect to the RPC node via web sockets
	client, err := wtCommon.GetConnection(config.L2WebsocketURL, config.Retries)
	if err != nil {
		wtCommon.Fatal(err)
	}

	// channel to recieve the headers of new blocks added to the chain
	headers := make(chan *types.Header)

	// resubscribe upton error automatically after 2 seconds
	sub := event.ResubscribeErr(2*time.Second, func(ctx context.Context, subErr error) (event.Subscription, error) {
		if subErr != nil {
			wtCommon.Warning("Subscription failed with folowing error")
			wtCommon.Error(subErr)
			wtCommon.Success("Auto-resubscribe is enabled, attempting resubscription...")
		}
		wtCommon.Info("Subscribing to Latest block heads on L2")

		eth_sub, err := client.SubscribeNewHead(context.Background(), headers)
		if err == nil {
			wtCommon.Success("Subscribed to the L2 blockchain head successfully!")
			wtCommon.Info("Waiting for block(s) ...")
		} else {
			wtCommon.Warning("Failed to subscribe to L2 Head")
			wtCommon.Error(err)
		}
		return eth_sub, err
	})

	return sub, headers, nil

}
