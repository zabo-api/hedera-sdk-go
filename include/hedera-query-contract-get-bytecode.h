#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-transaction-id.h"
#include "hedera-query.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__contract_get_bytecode__new(
    HederaClient*,
    HederaContractId contract
);

extern HederaError hedera_query__contract_get_bytecode__get(HederaQuery*, uint8_t**);

#ifdef __cplusplus
}
#endif
