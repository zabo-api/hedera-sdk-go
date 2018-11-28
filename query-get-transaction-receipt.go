package hedera

// #include "hedera-query-get-transaction-receipt.h"
import "C"

type QueryGetTransactionReceipt struct {
	Query
}

type TransactionReceipt struct {
	Status    uint8
	AccountID *AccountID
	// unsupported: ContractID *C.HederaContractId
	// unsupported: FileID *C.HederaFileId
}

const (
	TransactionStatus_UNKNOWN		uint8			= 0
	TransactionStatus_SUCCESS 		uint8			= 1
	TransactionStatus_FAIL_INVALID	uint8			= 2
	TransactionStatus_FAIL_FEE		uint8			= 3
	TransactionStatus_FAIL_BALANCE	uint8			= 4
)

var TransactionStatus_name = map[uint8]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "FAIL_INVALID",
	3: "FAIL_FEE",
	4: "FAIL_BALANCE",
}

var TransactionStatus_value = map[string]uint8{
	"UNKNOWN":      0,
	"SUCCESS":      1,
	"FAIL_INVALID": 2,
	"FAIL_FEE":     3,
	"FAIL_BALANCE": 4,
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

	receipt := TransactionReceipt{Status: uint8(answer.status)}

	if answer.account_id != nil {
		accountID := accountIDFromC(*answer.account_id)
		receipt.AccountID = &accountID
	}

	return receipt, nil
}
