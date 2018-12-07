package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"

type query struct {
	inner *C.HederaQuery
}

func (q query) Cost() (uint64, error) {
	var cost C.uint64_t
	res := C.hedera_query_cost(q.inner, &cost)
	if res != 0 {
		return 0, hederaLastError()
	}

	return uint64(cost), nil
}
