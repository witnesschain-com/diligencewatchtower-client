package interfaces

import (
	"fmt"

	datatypes "github.com/diligencewatchtower-client/coordinator/Datatypes"
	"github.com/diligencewatchtower-client/opchain"
	"github.com/diligencewatchtower-client/watcher"
)

func TraceTransaction(transactionHash string, deps datatypes.TracerDependencies) datatypes.TraceTxnResult {

	result := watcher.GetProofOfInclusion(
		[]string{transactionHash},
		&deps.Config,
		deps.Cache,
	)
	transactionReceipt := result[0].Receipt

	status := opchain.FetchTxnStatusFromReceipt(transactionReceipt, &deps.Config)

	res := datatypes.TraceTxnResult{
		Receipt: transactionReceipt,
		Status:  fmt.Sprint(status),
	}
	return res
}
