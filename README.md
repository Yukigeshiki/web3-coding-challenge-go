# web3-coding-challenge-go

Create an API in Go which allows a consumer to query / filter transfers on a specific ERC20 smart contract. It should be able to retrieve
- Transfers from a specific address
- Transfers to a specific address
- Transfers between 2 addresses
- Transfers above or below a certain value threshold

Other requirements:
- The api should give responses quickly, and not be forced to iterate through a contract's event history on every request
- The api should become aware of new transactions in real time (as they confirm)

### Solution Assumptions

- The start block for the ERC20 contract is known
- The service is only responsible for a single ERC20 contract address
- A longer start-up time is preferable to data inconsistency

### Example .env file

```.dotenv
PORT=8000
ETH_CLIENT_URL_HTTP=https://mainnet.infura.io/v3/<your-api-key>
ETH_CLIENT_URL_WS=wss://mainnet.infura.io/ws/v3/<your-api-key>
ERC20_CONTRACT_ADDRESS=<your-contract-address>
ERC20_CONTRACT_START_BLOCK=<start-block>
```

### Still needed

- Tests