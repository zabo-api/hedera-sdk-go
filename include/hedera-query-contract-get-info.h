#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-query.h"
#include "hedera-array.h"
#include "hedera-contract-info.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__contract_get_info__new(HederaClient*, HederaContractId id);

extern HederaError hedera_query__contract_get_info__get(HederaQuery*, HederaContractInfo**);

#ifdef __cplusplus
}
#endif
