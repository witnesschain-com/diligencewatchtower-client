package common

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetConnection(
	rawurl string,
	retries int,
) (*ethclient.Client, error) {
	Client, err := ethclient.Dial(rawurl)

	if err != nil {
		for connFailures := 1; connFailures < retries; connFailures++ {
			Client, err = ethclient.Dial(rawurl)
			if err == nil {
				break
			} else {
				Warning(err)
				Warning("Retry attempt " + strconv.Itoa(connFailures))
				time.Sleep(2 * time.Second)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return Client, nil
}

func GetRPCConnection(
	rpcURL string,
	retries int,
) (*rpc.Client, error) {
	Client, err := rpc.Dial(rpcURL)
	if err != nil {
		for connFailures := 1; connFailures < retries; connFailures++ {
			Client, err = rpc.Dial(rpcURL)
			if err == nil {
				break
			} else {
				Warning(err)
				Warning("Retry attempt " + strconv.Itoa(connFailures))
				time.Sleep(2 * time.Second)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return Client, nil
}
