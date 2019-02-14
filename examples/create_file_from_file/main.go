package main

import (
	"fmt"
	"io/ioutil"
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

	// load example file
	data, err := ioutil.ReadFile("../files/Hedera.txt")
	if err != nil {
		fmt.Println("File read error", err)
		return
	}
	//
	// Send transaction to create a file
	//

	nodeAccountID := hedera.AccountID{Account: 3}
	operatorAccountID := hedera.AccountID{Account: 2}

	expiryTime := time.Now().AddDate(0, 0, 1) // add one day

	response, err := client.CreateFile().
		Key(public).
		ExpiresAt(expiryTime).
		Contents(data).
		Operator(operatorAccountID).
		Node(nodeAccountID).
		Memo("[test] hedera-sdk-go file-create").
		Sign(operatorSecret).
		Sign(operatorSecret).
		Execute()

	if err != nil {
		panic(err)
	}

	transactionID := response.ID
	fmt.Printf("creating file; transaction = %v\n", transactionID)

	//
	// Get receipt to prove we created it ok
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

	fmt.Printf("file = %v\n", *receipt.FileID)
}
