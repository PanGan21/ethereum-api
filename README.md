# ethereum-api
Ethereum api using go lang

## Run it
```go run main.go```<br/>
or if you want to run it with [ganache](https://github.com/trufflesuite/ganache-cli) use<br/>
```docker-compose up --build```

## Api endpoints
- Fetch the latest block: ```http://localhost:8080/api/eth/latest-block```
- Send some ethers using Curl: ```curl -d '{"privKey":"<private-key>", "to":"<account-to-send>", "amount":<amount>}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/eth/send-eth```
- Get the transaction information: ```http://localhost:8080/api/eth/get-tx?hash=<tx-hash-from-response>```
- Check your balance: ```http://localhost:8080/api/eth/get-balance?address=<the-recipient-address>```
