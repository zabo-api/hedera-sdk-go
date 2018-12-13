package hedera

// #include <stdlib.h>
// #include "hedera-transaction-id.h"
import "C"
import (
	"time"
	"unsafe"
)

type TransactionID struct {
	// Deprecated: The response is already a TransactionID
	ID *TransactionID

	AccountID             AccountID
	TransactionValidStart time.Time
}

func TransactionIDFromString(s string) (TransactionID, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var id C.HederaTransactionId
	res := C.hedera_transaction_id_from_str(cStr, &id)
	if res != 0 {
		return TransactionID{}, hederaLastError()
	}

	return goTransactionID(id), nil
}

func (id TransactionID) String() string {
	cID := cTransactionID(id)
	return hederaString(C.hedera_transaction_id_to_str(&cID))
}

func cTransactionID(id TransactionID) C.HederaTransactionId {
	return C.HederaTransactionId{
		account_id: cAccountID(id.AccountID),
		transaction_valid_start: C.HederaTimestamp{
			seconds: C.int64_t(id.TransactionValidStart.Unix()),
			nanos:   C.uint32_t(id.TransactionValidStart.Nanosecond()),
		},
	}
}

func goTransactionID(id C.HederaTransactionId) TransactionID {
	goID := TransactionID{
		AccountID: goAccountID(id.account_id),
		TransactionValidStart: time.Unix(
			int64(id.transaction_valid_start.seconds),
			int64(id.transaction_valid_start.nanos)),
	}

	goID.ID = &goID

	return goID
}
