#pragma once

#include <stdint.h>
#include "hedera-crypto.h"
#include "hedera-query.h"
#include "hedera-entity.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__get_by_key__new(
    HederaClient*,
    HederaPublicKey key
);

extern HederaError hedera_query__get_by_key__get(
    HederaQuery*, 
    HederaEntity**
);

#ifdef __cplusplus
}
#endif