package main

import (
	"fmt"
	"os"

	"github.com/zabo-api/hedera-sdk-go"
)

func main() {
	// Target account to get the balance for
	accountID := hedera.AccountID{Account: 2}

	client, err := hedera.Dial("testnet.hedera.com:50003")
	if err != nil {
		panic(err)
	}

	client.SetNode(hedera.AccountID{Account: 3})
	client.SetOperator(accountID, func() hedera.SecretKey {
		operatorSecret, err := hedera.SecretKeyFromString(os.Getenv("OPERATOR_SECRET"))
		if err != nil {
			panic(err)
		}

		return operatorSecret
	})

	defer client.Close()

	// Get the _answer_ for the query of getting the account balance
	balance, err := client.Account(accountID).Balance().Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("balance = %v tinybars\n", balance)
	fmt.Printf("balance = %.5f hbars\n", float64(balance)/100000000.0)
}
