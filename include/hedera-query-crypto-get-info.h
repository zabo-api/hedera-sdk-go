#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-query.h"
#include "hedera-array.h"
#include "hedera-account-info.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__crypto_get_info__new(
    HederaClient*,
    HederaAccountId account_id
);

extern HederaError hedera_query__crypto_get_info__get(
    HederaQuery*, 
    HederaAccountInfo**
);

#ifdef __cplusplus
}
#endif