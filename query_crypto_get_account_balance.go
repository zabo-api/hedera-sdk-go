package hedera

// #include "hedera.h"
import "C"
import "github.com/markbates/oncer"

type QueryCryptoGetAccountBalance struct {
	query
}

func newQueryCryptoGetAccountBalance(client Client, id AccountID) QueryCryptoGetAccountBalance {
	return QueryCryptoGetAccountBalance{
		query{C.hedera_query__crypto_get_account_balance__new(client.inner, cAccountID(id))},
	}
}

// Deprecated: Use Query.Get() instead
func (query QueryCryptoGetAccountBalance) Answer() (uint64, error) {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Query.Answer()",
		"Use Query.Get() instead.")

	return query.Get()
}

func (query QueryCryptoGetAccountBalance) Get() (uint64, error) {
	var balance C.uint64_t
	res := C.hedera_query__crypto_get_account_balance__get(query.inner, &balance)
	if res != 0 {
		return 0, hederaLastError()
	}

	return uint64(balance), nil
}
