package ethclient

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c EthClient) FilterLogs(query ethereum.FilterQuery) []types.Log {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var logs []types.Log
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "filterLogs")
		logs, err = c.ethClient.FilterLogs(c.ctx, query)
	}
	return logs
}

func (c EthClient) ChainID() *big.Int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var chainID *big.Int
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "chainID")
		chainID, err = c.ethClient.ChainID(c.ctx)
	}
	return chainID
}

func (c EthClient) BlockNumber() uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var height uint64
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "blockNumber")
		height, err = c.ethClient.BlockNumber(c.ctx)
	}
	return height
}

func (c EthClient) SendTransaction(tx *types.Transaction) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "sendTransaction")
		err = c.ethClient.SendTransaction(c.ctx, tx)
	}
}

func (c EthClient) NetworkID() *big.Int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var networkID *big.Int
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "networkID")
		networkID, err = c.ethClient.NetworkID(c.ctx)
	}
	return networkID
}

func (c EthClient) PendingNonceAt(account common.Address) uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var nonce uint64
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "pendingNonceAt")
		nonce, err = c.ethClient.PendingNonceAt(c.ctx, account)
	}
	return nonce
}

func (c EthClient) SuggestGasPrice() *big.Int {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var price *big.Int
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "suggestGasPrice")
		price, err = c.ethClient.SuggestGasPrice(c.ctx)
	}
	return price
}

func (c EthClient) TransactionReceipt(txHash common.Hash) *types.Receipt {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var receipt *types.Receipt
	err := fmt.Errorf("error")
	retry := 0
	for err != nil {
		retry = c.retryCall(retry, "transactionReceipt")
		receipt, err = c.ethClient.TransactionReceipt(c.ctx, txHash)
	}
	return receipt
}
