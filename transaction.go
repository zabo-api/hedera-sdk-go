package hedera

// #include "hedera-transaction.h"
import "C"
import "unsafe"

type Transaction struct {
	inner *C.HederaTransaction
}

type TransactionResponse struct {
	ID TransactionID
}

func (tx Transaction) Operator(id AccountID) Transaction {
	C.hedera_transaction_set_operator(tx.inner, id.c())
	return tx
}

func (tx Transaction) Node(id AccountID) Transaction {
	C.hedera_transaction_set_node(tx.inner, id.c())
	return tx
}

func (tx Transaction) Memo(memo string) Transaction {
	C.hedera_transaction_set_memo(tx.inner, C.CString(memo))
	return tx
}

func (tx Transaction) Sign(key SecretKey) Transaction {
	C.hedera_transaction_sign(tx.inner, key.inner)
	return tx
}

func (tx Transaction) Execute() (TransactionResponse, error) {
	var res C.HederaTransactionResponse
	err := C.hedera_transaction_execute(tx.inner, &res)

	return *((*TransactionResponse)(unsafe.Pointer(&res))), hederaError(err)
}
