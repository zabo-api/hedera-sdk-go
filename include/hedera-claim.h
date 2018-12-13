#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-crypto.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    HederaAccountId account;
    
    uint8_t* hash;
    size_t hash_len;
    
    HederaPublicKey* keys;
    size_t keys_len;
} HederaClaim;

#ifdef __cplusplus
}
#endif
