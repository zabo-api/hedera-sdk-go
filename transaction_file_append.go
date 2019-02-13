package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import "unsafe"

type TransactionFileAppend struct {
	transaction
}

func newTransactionFileAppend(client *Client, id FileID, contents []byte) TransactionFileAppend {
	cLength := C.ulong(len(contents))
	cContents := (*C.uchar)(unsafe.Pointer(&contents[0]))
	return TransactionFileAppend{transaction{
		C.hedera_transaction__file_append__new(client.inner, cFileID(id), cContents, cLength)}}
}

func (tx TransactionFileAppend) Operator(id AccountID) TransactionFileAppend {
	return TransactionFileAppend{tx.transaction.Operator(id)}
}

func (tx TransactionFileAppend) Node(id AccountID) TransactionFileAppend {
	return TransactionFileAppend{tx.transaction.Node(id)}
}

func (tx TransactionFileAppend) Memo(memo string) TransactionFileAppend {
	return TransactionFileAppend{tx.transaction.Memo(memo)}
}
