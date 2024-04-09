package external

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
	wtCommon "github.com/diligencewatchtower-client/common"
	"github.com/prometheus/client_golang/prometheus"
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
