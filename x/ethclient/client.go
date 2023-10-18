package ethclient

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient struct {
	ethClient *ethclient.Client
	mutex     *sync.Mutex
	wsURL     string
	maxRetry  int
	ctx       context.Context
}

func NewClient(ctx context.Context, wsURL string, maxRetry int) *EthClient {
	client := &EthClient{
		ethClient: nil,
		mutex:     &sync.Mutex{},
		wsURL:     wsURL,
		maxRetry:  maxRetry,
		ctx:       ctx,
	}
	client.ethClient = client.GetEthereumClient()
	return client
}

func (c EthClient) GetEthereumClient() *ethclient.Client {
	var client *ethclient.Client
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "ethclient")
		client, err = ethclient.Dial(c.wsURL)
	}

	return client
}
