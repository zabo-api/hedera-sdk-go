package main

import (
	"fmt"
	"github.com/mehcode/hedera-sdk-go"
)

func main() {
	client := hedera.Dial("testnet.hedera.com:50001")
	defer client.Close()

	accountID := hedera.NewAccountID(0, 0, 2)
	balance, err := client.GetAccountBalance(accountID).Answer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("balance = %v\n", balance)
}
