package hedera

// #include "hedera-account-id.h"
import "C"
import (
	"unsafe"
)

type AccountID struct {
	Realm   int64 `json:"realm"`
	Shard   int64 `json:"shard"`
	Account int64 `json:"account"`
}

// Parse an account ID from the string.
// Expects a string of the form: {realm}:{shard}:{account}
func AccountIDFromString(s string) (AccountID, error) {
	var accountID C.HederaAccountId
	err := C.hedera_account_id_from_str(C.CString(s), &accountID)
	if err != 0 {
		return AccountID{}, hederaError(err)
	}

	return accountIDFromC(accountID), nil
}

// Cast an [AccountId] from the C SDK
func accountIDFromC(id C.HederaAccountId) AccountID {
	return *((*AccountID)(unsafe.Pointer(&id)))
}

// Cast an [AccountId] to the C SDK
func (id AccountID) c() C.HederaAccountId {
	return *(*C.HederaAccountId)(unsafe.Pointer(&id))
}
