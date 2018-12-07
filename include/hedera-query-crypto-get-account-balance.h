#pragma once

#include "hedera-id.h"
#include "hedera-query.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__crypto_get_account_balance__new(HederaClient*, HederaAccountId account);

extern HederaError hedera_query__crypto_get_account_balance__get(HederaQuery*, uint64_t*);

#ifdef __cplusplus
}
#endif
