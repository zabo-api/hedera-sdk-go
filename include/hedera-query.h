#pragma once

#include "hedera-error.h"
#include "hedera-client.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct HederaQuery HederaQuery;

extern HederaError hedera_query_cost(HederaQuery*, uint64_t*);

#ifdef __cplusplus
}
#endif
