package main

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
	"os"
	"time"
)

func main() {
	// Read and decode the operator secret key
	operatorAccountID := hedera.AccountID{Account: 2}
	operatorSecret, err := hedera.SecretKeyFromString(os.Getenv("OPERATOR_SECRET"))
	if err != nil {
		panic(err)
	}

	// Read and decode target account
	targetAccountID, err := hedera.AccountIDFromString(os.Getenv("TARGET"))
	if err != nil {
		panic(err)
	}

	//
	// Connect to Hedera
	//

	client, err := hedera.Dial("testnet.hedera.com:50003")
	if err != nil {
		panic(err)
	}

	client.SetNode(hedera.AccountID{Account: 3})
	client.SetOperator(operatorAccountID, func() hedera.SecretKey {
		return operatorSecret
	})

	defer client.Close()

	//
	// Get balance for target account
	//

	balance, err := client.Account(targetAccountID).Balance().Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("account balance = %v\n", balance)

	//
	// Transfer 100 cryptos to target
	//

	nodeAccountID := hedera.AccountID{Account: 3}
	response, err := client.TransferCrypto().
		// Move 100 out of operator account
		Transfer(operatorAccountID, -100).
		// And place in our new account
		Transfer(targetAccountID, 100).
		Operator(operatorAccountID).
		Node(nodeAccountID).
		Memo("[test] hedera-sdk-go v2").
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

	receipt, err := client.Transaction(*transactionID).Receipt().Get()
	if err != nil {
		panic(err)
	}

	if receipt.Status != hedera.StatusSuccess {
		panic(fmt.Errorf("transaction has a non-successful status: %v", receipt.Status.String()))
	}

	fmt.Printf("wait for 2s...\n")
	time.Sleep(2 * time.Second)

	//
	// Get balance for target account (again)
	//

	balance, err = client.Account(targetAccountID).Balance().Get()
	if err != nil {
		panic(err)
	}

	fmt.Printf("account balance = %v\n", balance)
}
