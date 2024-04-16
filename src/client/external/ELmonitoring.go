package external

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	"github.com/prometheus/client_golang/prometheus"
	wtCommon "github.com/witnesschain-com/diligencewatchtower-client/common"
)

func InitialiseELMonitoring() {

	logger, err := logging.NewZapLogger("production")
	if err != nil {
		wtCommon.Error(err)
	}

	reg := prometheus.NewRegistry()
	eigenMetrics := metrics.NewEigenMetrics("watchtower", ":9090", reg, logger)
	eigenMetrics.Start(context.Background(), reg)
}
