package hedera

// #include "hedera.h"
import "C"
import "github.com/markbates/oncer"

type QueryGetTransactionReceipt struct {
	query
}

func newQueryGetTransactionReceipt(client Client, id TransactionID) QueryGetTransactionReceipt {
	return QueryGetTransactionReceipt{
		query{C.hedera_query__get_transaction_receipt__new(client.inner, cTransactionID(id))},
	}
}

// Deprecated: Use Query.Get() instead
func (query QueryGetTransactionReceipt) Answer() (TransactionReceipt, error) {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Query.Answer()",
		"Use Query.Get() instead.")

	return query.Get()
}

func (query QueryGetTransactionReceipt) Get() (TransactionReceipt, error) {
	var out C.HederaTransactionReceipt
	res := C.hedera_query__get_transaction_receipt__get(query.inner, &out)
	if res != 0 {
		return TransactionReceipt{}, hederaLastError()
	}

	receipt := TransactionReceipt{Status: TransactionStatus(out.status)}

	if out.account_id != nil {
		accountID := goAccountID(*out.account_id)
		receipt.AccountID = &accountID
	}

	if out.contract_id != nil {
		contractID := goContractID(*out.contract_id)
		receipt.ContractID = &contractID
	}

	if out.file_id != nil {
		fileID := goFileID(*out.file_id)
		receipt.FileID = &fileID
	}

	return receipt, nil
}
