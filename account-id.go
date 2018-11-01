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

func NewAccountID(realm, shard, account int64) AccountID {
	return AccountID{Realm: realm, Shard: shard, Account: account}
}

func AccountIDFromString(s string) (AccountID, error) {
	var accountID C.HederaAccountId
	err := C.hedera_account_id_from_str(C.CString(s), &accountID)
	if err != 0 {
		return AccountID{}, hederaError(err)
	}

	return accountIDFromC(accountID), nil
}

func accountIDFromC(id C.HederaAccountId) AccountID {
	return *((*AccountID)(unsafe.Pointer(&id)))
}

func (id AccountID) c() C.HederaAccountId {
	return *(*C.HederaAccountId)(unsafe.Pointer(&id))
}
