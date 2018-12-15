#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-query.h"
#include "hedera-transaction-record.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__contract_get_records__new(
    HederaClient*,
    HederaContractId contract_id
);

extern HederaError hedera_query__contract_get_records__get(HederaQuery*, HederaArray**);

#ifdef __cplusplus
}
#endif
