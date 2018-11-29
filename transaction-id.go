package hedera

// #include <stdlib.h>
// #include "hedera-transaction-id.h"
import "C"
import (
	"time"
	"unsafe"
)

type TransactionID struct {
	AccountID             AccountID
	TransactionValidStart time.Time
}

func (id TransactionID) c() C.HederaTransactionId {
	return C.HederaTransactionId{
		account_id: id.AccountID.c(),
		transaction_valid_start: C.HederaTimestamp{
			seconds: C.int64_t(id.TransactionValidStart.Unix()),
			nanos:   C.int32_t(id.TransactionValidStart.Nanosecond()),
		},
	}
}

func (id TransactionID) String() string {
	bytes := C.hedera_transaction_id_to_str(id.c())
	defer C.free(unsafe.Pointer(bytes))

	return C.GoString(bytes)
}

func transactionIDFromC(id C.HederaTransactionId) TransactionID {
	return TransactionID{
		AccountID: accountIDFromC(id.account_id),
		TransactionValidStart: time.Unix(
			int64(id.transaction_valid_start.seconds),
			int64(id.transaction_valid_start.nanos)),
	}
}

func TransactionIDFromString(s string) (TransactionID, error) {
	var transactionID C.HederaTransactionId
	err := C.hedera_transaction_id_from_str(C.CString(s), &transactionID)
	if err != 0 {
		return TransactionID{}, hederaError(err)
	}

	return transactionIDFromC(transactionID), nil
}
