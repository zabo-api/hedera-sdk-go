package main

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	client, err := hedera.Dial("testnet.hedera.com:50001")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// Target account to get the balance for
	accountID := hedera.AccountID{Account: 2}

	// Get the _answer_ for the query of getting the account balance
	balance, err := client.Account(accountID).Balance().Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("balance = %v tinybars\n", balance)
	fmt.Printf("balance = %.5f hbars\n", float64(balance)/100000000.0)
}
