package main

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
	"os"
	"time"
)

func main() {
	// Read and decode the operator secret key
	operatorSecret, err := hedera.SecretKeyFromString(os.Getenv("OPERATOR_SECRET"))
	if err != nil {
		panic(err)
	}

	// Read and decode target account
	targetAccountId, err := hedera.AccountIDFromString(os.Getenv("TARGET"))
	if err != nil {
		panic(err)
	}

	//
	// Connect to Hedera
	//

	client := hedera.Dial("testnet.hedera.com:50001")
	// TODO: client.SetRetryOnFailure(0) // default: 5
	defer client.Close()

	//
	// Get balance for target account
	//

	balance, err := client.GetAccountBalance(targetAccountId).Answer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("account balance = %v\n", balance)

	//
	// Transfer 100 cryptos to target
	//

	nodeAccountId := hedera.NewAccountID(0, 0, 3)
	operatorAccountID := hedera.NewAccountID(0, 0, 2)
	response, err := client.CryptoTransfer().
		Operator(operatorAccountID).
		Node(nodeAccountId).
		Memo("[test] hedera-sdk-go v2").
		// Move 100 out of operator account
		Transfer(operatorAccountID, -100).
		// And place in our new account
		Transfer(targetAccountId, 100).
		Sign(operatorSecret). // Sign it once as operator
		Sign(operatorSecret). // And again as sender
		Execute()

	if err != nil {
		panic(err)
	}

	transactionID := response.ID
	fmt.Printf("transferred; transaction = %v\n", transactionID)

	//
	// Get receipt to prove we sent ok
	//

	fmt.Printf("wait for 2s...\n")
	time.Sleep(2 * time.Second)

	_, err = client.GetTransactionReceipt(transactionID).Answer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("wait for 2s...\n")
	time.Sleep(2 * time.Second)

	//
	// Get balance for target account (again)
	//

	balance, err = client.GetAccountBalance(targetAccountId).Answer()
	if err != nil {
		panic(err)
	}

	fmt.Printf("account balance = %v\n", balance)
}
