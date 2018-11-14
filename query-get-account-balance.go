package hedera

// #include "hedera-query-get-account-balance.h"
import "C"

type QueryGetAccountBalance struct {
	Query
}

func newQueryGetAccountBalance(client *Client, account AccountID) QueryGetAccountBalance {
	return QueryGetAccountBalance{
		Query{C.hedera_query__get_account_balance__new(client.inner, account.c())}}
}

func (query QueryGetAccountBalance) Answer() (uint64, error) {
	var answer C.uint64_t
	err := C.hedera_query__get_account_balance__answer(query.inner, &answer)
	if err != 0 {
		return 0, hederaError(err)
	}

	return uint64(answer), nil
}
