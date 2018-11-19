package hedera

// #include "hedera-account-id.h"
import "C"
import (
	"github.com/markbates/oncer"
	"unsafe"
)

type AccountID struct {
	Realm   int64 `json:"realm"`
	Shard   int64 `json:"shard"`
	Account int64 `json:"account"`
}

func NewAccountID(realm, shard, account int64) AccountID {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#NewAccountID",
		"Use github.com/hashgraph/hedera-sdk-go#AccountID instead.")

	return AccountID{Realm: realm, Shard: shard, Account: account}
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
