package main

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
	"time"
)

func main() {
	client, err := hedera.Dial("testnet.hedera.com:50001")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// Target account to get the balance for
	accountID := hedera.AccountID{Account: 2}

	// Get the _cost_ or transaction fee for the query of getting the account balance
	cost, err := client.Account(accountID).Balance().Cost()
	if err != nil {
		panic(err)
	}

	// Wait 1s between queries (testnet transaction limitation)
	time.Sleep(1 * time.Second)

	// Get the _answer_ for the query of getting the account balance
	balance, err := client.Account(accountID).Balance().Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("cost    = %v tinybars\n", cost)
	fmt.Printf("balance = %v tinybars\n", balance)
	fmt.Printf("balance = %.5f hbars\n", float64(balance)/100000000.0)
}
