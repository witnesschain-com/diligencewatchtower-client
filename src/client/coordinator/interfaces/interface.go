package interfaces

import (
	"fmt"

	"github.com/witnesschain-com/diligencewatchtower-client/opchain"
	"github.com/witnesschain-com/diligencewatchtower-client/watcher"
	datatypes "github.com/witnesschain-com/diligencewatchtower-client/watchtower/coordinator/Datatypes"
)

func TraceTransaction(transactionHash string, deps datatypes.TracerDependencies) datatypes.TraceTxnResult {

	result := watcher.GetProofOfInclusion(
		[]string{transactionHash},
		&deps.Config,
		deps.Vault,
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
