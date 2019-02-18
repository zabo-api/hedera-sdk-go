package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	// Read and decode the operator secret key
	operatorSecret, err := hedera.SecretKeyFromString(os.Getenv("OPERATOR_SECRET"))
	if err != nil {
		panic(err)
	}

	// Generate a new keypair for the new account
	public := operatorSecret.Public()

	fmt.Printf("public = %v\n", public)

	//
	// Connect to Hedera
	//

	client, err := hedera.Dial("testnet.hedera.com:50003")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	//
	// Send transaction to create a file
	//

	nodeAccountID := hedera.AccountID{Account: 3}
	operatorAccountID := hedera.AccountID{Account: 2}

	expiryTime := time.Now().AddDate(0, 0, 1) // add one day

	client.SetNode(nodeAccountID)
	client.SetOperator(operatorAccountID, func() hedera.SecretKey {
		operatorSecret, err := hedera.SecretKeyFromString(os.Getenv("OPERATOR_SECRET"))
		if err != nil {
			panic(err)
		}

		return operatorSecret
	})

	response, err := client.CreateFile().
		Key(public).
		ExpiresAt(expiryTime).
		Contents([]byte("Hedera Hashgraph is great")).
		// Operator(operatorAccountID).
		// Node(nodeAccountID).
		Memo("[test] hedera-sdk-go file-create").
		GenerateRecord(true).
		Sign(operatorSecret).
		// Sign(operatorSecret).
		Execute()

	if err != nil {
		panic(err)
	}

	transactionID := response.ID
	fmt.Printf("creating file; transaction = %v\n", transactionID)

	//
	// Get receipt to prove we created it ok
	//

	fmt.Printf("wait for 5s...\n")
	time.Sleep(5 * time.Second)

	fmt.Printf("getting receipt; transaction = %v\n", transactionID)
	receipt, err := client.Transaction(*transactionID).Receipt().Get()
	if err != nil {
		panic(err)
	}

	if receipt.Status != hedera.StatusSuccess {
		panic(fmt.Errorf("transaction has a non-successful status: %v", receipt.Status.String()))
	}

	fmt.Printf("file = %v\n", *receipt.FileID)

	fmt.Printf("wait for 5s...\n")
	time.Sleep(5 * time.Second)

	record, err := client.Transaction(*transactionID).Record().
		Get()

	if err != nil {
		panic(err)
	}

	fmt.Printf("hash = %v\n", record)

}
