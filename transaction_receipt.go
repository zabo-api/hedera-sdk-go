package hedera

type TransactionReceipt struct {
	Status     TransactionStatus
	AccountID  *AccountID
	ContractID *ContractID
	FileID     *FileID
}

type TransactionStatus uint8

const (
	TransactionStatusUnkown      TransactionStatus = 0
	TransactionStatusSuccess     TransactionStatus = 1
	TransactionStatusFailInvalid TransactionStatus = 2
	TransactionStatusFailFee     TransactionStatus = 3
	TransactionStatusFailBalance TransactionStatus = 4
)

var transactionStatusText = map[TransactionStatus]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAIL_INVALID",
	3: "FAIL_FEE",
	4: "FAIL_BALANCE",
}

func (status TransactionStatus) String() string {
	return transactionStatusText[status]
}
