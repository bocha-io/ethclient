package ethclient

import (
	"fmt"
	"time"

	"github.com/bocha-io/logger"
)

func (c *EthClient) retryCall(retry int, functionName string) int {
	if retry == 0 {
		return 1
	}

	if retry < c.maxRetry {
		logger.LogError(fmt.Sprintf("[indexer] retrying (%s) (%d)", functionName, retry))
		time.Sleep(time.Duration(retry) * (time.Second))
		retry++
		// Recreate the connection
		c.ethClient = c.GetEthereumClient()
		return retry
	}

	logger.LogError(fmt.Sprintf("[indexer] rpc down (%s)", functionName))
	panic(fmt.Sprintf("[indexer] rpc down (%s)", functionName))
}
