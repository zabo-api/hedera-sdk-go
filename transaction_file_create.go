package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import "unsafe"
import "time"

type TransactionFileCreate struct {
	transaction
}

func newTransactionFileCreate(client *Client) TransactionFileCreate {
	return TransactionFileCreate{transaction{
		C.hedera_transaction__file_create__new(client.inner)}}
}

func (tx TransactionFileCreate) Key(public PublicKey) TransactionFileCreate {
	C.hedera_transaction__file_create__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionFileCreate) Contents(contents []byte) TransactionFileCreate {
	cLength := C.ulong(len(contents))
	cContents := (*C.uchar)(unsafe.Pointer(&contents[0]))
	C.hedera_transaction__file_create__set_contents(tx.inner, cContents, cLength)
	return tx
}

func (tx TransactionFileCreate) ExpiresAt(expiresAt time.Time) TransactionFileCreate {

	cExpiration := C.HederaTimestamp{
		seconds: C.int64_t(expiresAt.Unix()),
		nanos:   C.uint32_t(expiresAt.Nanosecond()),
	}

	C.hedera_transaction__file_create__set_expires_at(tx.inner, cExpiration)
	return tx
}

func (tx TransactionFileCreate) Operator(id AccountID) TransactionFileCreate {
	return TransactionFileCreate{tx.transaction.Operator(id)}
}

func (tx TransactionFileCreate) Node(id AccountID) TransactionFileCreate {
	return TransactionFileCreate{tx.transaction.Node(id)}
}

func (tx TransactionFileCreate) Memo(memo string) TransactionFileCreate {
	return TransactionFileCreate{tx.transaction.Memo(memo)}
}
