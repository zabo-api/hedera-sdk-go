package hedera

// #include "hedera-query-get-account-balance.h"
import "C"

type QueryGetAccountBalance struct {
	inner *C.HederaQuery
}

func newQueryGetAccountBalance(client *Client, account AccountID) QueryGetAccountBalance {
	return QueryGetAccountBalance{
		C.hedera_query__get_account_balance__new(client.inner, account.c())}
}

func (query QueryGetAccountBalance) Cost() (uint64, error) {
	var cost C.uint64_t
	err := C.hedera_query__get_account_balance__cost(query.inner, &cost)
	if err != 0 {
		return 0, hederaError(err)
	}

	return uint64(cost), nil
}

func (query QueryGetAccountBalance) Answer() (uint64, error) {
	var answer C.uint64_t
	err := C.hedera_query__get_account_balance__answer(query.inner, &answer)
	if err != 0 {
		return 0, hederaError(err)
	}

	return uint64(answer), nil
}
