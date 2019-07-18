package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import "unsafe"

type transaction struct {
	inner *C.HederaTransaction
}

type RawTransaction struct {
	inner *C.HederaTransaction
}

func (tx transaction) Operator(id AccountID) transaction {
	C.hedera_transaction_set_operator(tx.inner, cAccountID(id))
	return tx
}

func (tx transaction) Node(id AccountID) transaction {
	C.hedera_transaction_set_node(tx.inner, cAccountID(id))
	return tx
}

func (tx transaction) Memo(memo string) transaction {
	cStr := C.CString(memo)
	defer C.free(unsafe.Pointer(cStr))

	C.hedera_transaction_set_memo(tx.inner, cStr)
	return tx
}

func (tx transaction) Fee(fee uint64) transaction {
	C.hedera_transaction_set_fee(tx.inner, C.ulong(fee))
	return tx
}

func (tx transaction) Sign(key SecretKey) RawTransaction {
	return RawTransaction{tx.inner}.Sign(key)
}

func (tx transaction) Execute() (TransactionID, error) {
	return RawTransaction{tx.inner}.Execute()
}

func (tx RawTransaction) Sign(key SecretKey) RawTransaction {
	C.hedera_transaction_sign(tx.inner, &key.inner)
	return tx
}

func (tx RawTransaction) Execute() (TransactionID, error) {
	var id C.HederaTransactionId
	res := C.hedera_transaction_execute(tx.inner, &id)
	if res != 0 {
		return TransactionID{}, hederaLastError()
	}

	return goTransactionID(id), nil
}
