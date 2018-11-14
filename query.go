package hedera

// #include "hedera-query.h"
import "C"

type Query struct {
	inner *C.HederaQuery
}

func (query Query) Cost() (uint64, error) {
	var cost C.uint64_t
	err := C.hedera_query_cost(query.inner, &cost)
	if err != 0 {
		return 0, hederaError(err)
	}

	return uint64(cost), nil
}
