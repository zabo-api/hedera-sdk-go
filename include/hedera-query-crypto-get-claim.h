#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-query.h"
#include "hedera-claim.h"
#include "hedera-array.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__crypto_get_claim__new(
    HederaClient*,
    HederaAccountId account_id,
    HederaArray hash,
);

extern HederaError hedera_query__crypto_get_claim__get(HederaQuery*, HederaClaim**);

#ifdef __cplusplus
}
#endif