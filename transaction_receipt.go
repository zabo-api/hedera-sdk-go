package hedera

type TransactionReceipt struct {
	Status     TransactionResponse
	AccountID  *AccountID
	ContractID *ContractID
	FileID     *FileID
}

func (status TransactionResponse) String() string {
	return transactionResponseText[status]
}
