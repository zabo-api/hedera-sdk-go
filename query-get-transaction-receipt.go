package hedera

// #include "hedera-query-get-transaction-receipt.h"
import "C"

type QueryGetTransactionReceipt struct {
	Query
}

type TransactionStatus uint8

const (
	TransactionStatusUnkown			TransactionStatus		= 0
	TransactionStatusSuccess 		TransactionStatus		= 1
	TransactionStatusFailInvalid	TransactionStatus		= 2
	TransactionStatusFailFee		TransactionStatus		= 3
	TransactionStatusFailBalance	TransactionStatus		= 4
)

var transactionStatusText = map[TransactionStatus]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAIL_INVALID",
	3: "FAIL_FEE",
	4: "FAIL_BALANCE",
}

type TransactionReceipt struct {
	Status    TransactionStatus
	AccountID *AccountID
	// unsupported: ContractID *C.HederaContractId
	// unsupported: FileID *C.HederaFileId
}

func (status TransactionStatus) String() string {
	return transactionStatusText[status]
}

func newQueryGetTransactionReceipt(client *Client, transactionID TransactionID) QueryGetTransactionReceipt {
	return QueryGetTransactionReceipt{
		Query{C.hedera_query__get_transaction_receipt__new(client.inner, transactionID.c())}}
}

func (query QueryGetTransactionReceipt) Answer() (TransactionReceipt, error) {
	var answer C.HederaQueryGetTransactionReceiptAnswer
	err := C.hedera_query__get_transaction_receipt__answer(query.inner, &answer)
	if err != 0 {
		return TransactionReceipt{}, hederaError(err)
	}

	receipt := TransactionReceipt{Status: TransactionStatus(answer.status)}

	if answer.account_id != nil {
		accountID := accountIDFromC(*answer.account_id)
		receipt.AccountID = &accountID
	}

	return receipt, nil
}
