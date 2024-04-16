package contractutils

import (
	"context"
	"time"

	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitForTransactionReceipt(
	config *wtCommon.SimplifiedConfig,
	txn *types.Transaction,
	L1Client *ethclient.Client,
) (*types.Receipt, error) {
	wtCommon.Info("Waiting for transaction receipt ...")

	// wait upto `ReceiptTimeout` seconds for txn receipt
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ReceiptTimeout)*(time.Second))

	// releases resources if slowOperation completes before timeout elapses
	defer cancel()

	receipt, err := bind.WaitMined(ctx, L1Client, txn)
	if err != nil {
		wtCommon.Error(err)
	} else if receipt.Status == 1 {
		wtCommon.Info("Transaction executed successfully, logs are ...")
		wtCommon.Success(receipt.Logs)
	} else if receipt.Status == 0 {
		wtCommon.Error("Transaction submitted successfully but failed to execute!")
	}

	return receipt, err
}
