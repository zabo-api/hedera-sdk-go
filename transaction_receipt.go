package hedera

type TransactionReceipt struct {
	Status     TransactionStatus
	AccountID  *AccountID
	ContractID *ContractID
	FileID     *FileID
}

func (status TransactionStatus) String() string {
	return transactionStatusText[status]
}
