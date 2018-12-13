package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"

type query struct {
	inner *C.HederaQuery
}
