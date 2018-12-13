package main

import (
	"fmt"
	"github.com/hashgraph/hedera-sdk-go"
)

func main() {
	secret, mnemonic := hedera.GenerateSecretKey()
	fmt.Printf("secret   = %v\n", secret)
	fmt.Printf("mnemonic = %v\n", mnemonic)

	public := secret.Public()
	fmt.Printf("public   = %v\n", public)
}
