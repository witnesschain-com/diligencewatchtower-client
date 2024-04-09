package external

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	wtCommon "github.com/diligencewatchtower-client/common"
)

func GetELLogger() logging.Logger {
	logger, err := logging.NewZapLogger("development")
	if err != nil {
		wtCommon.Error(err)
	}
	return logger
}
