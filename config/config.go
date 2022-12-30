package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ContractStartBlock uint64
	ERC20ContractAddr  string
	EthClientURLHTTP   string
	EthClientURLWS     string
	Port               string
)

// init loads a .env file if the environment is local, then assigns environment variables to config values
func init() {

	if os.Getenv("ENV") == "" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ContractStartBlock, _ = strconv.ParseUint(os.Getenv("ERC20_CONTRACT_START_BLOCK"), 10, 64)
	ERC20ContractAddr = os.Getenv("ERC20_CONTRACT_ADDRESS")
	EthClientURLHTTP = os.Getenv("ETH_CLIENT_URL_HTTP")
	EthClientURLWS = os.Getenv("ETH_CLIENT_URL_WS")
	Port = os.Getenv("PORT")
}
