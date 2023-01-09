package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"robothouse.io/web3-coding-challenge/lib/ethereum/contracts"
)

func GetContractInstance(url, contractAddr string) (*contracts.ERC20, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	tokenAddress := common.HexToAddress(contractAddr)
	return contracts.NewERC20(tokenAddress, client)
}
