#pragma once

#include <stdint.h>
#include "hedera-transaction.h"
#include "hedera-claim.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__crypto_update__new(
    HederaClient*, 
    HederaAccountId account_id
);

extern void hedera_transaction__crypto_update__set_key(
    HederaTransaction*, 
    HederaPublicKey key
);

extern void hedera_transaction__crypto_update__set_proxy_account(
    HederaTransaction*, 
    HederaAccountId proxy_account
);


extern void hedera_transaction__crypto_update__set_proxy_fraction(
    HederaTransaction*, 
    int32_t proxy_fraction
);

extern void hedera_transaction__crypto_update__set_send_record_threshold(
    HederaTransaction*, 
    uint64_t send_record_threshold
);

extern void hedera_transaction__crypto_update__set_receive_record_threshold(
    HederaTransaction*, 
    uint64_t receive_record_threshold
);

#ifdef __cplusplus
}
#endif
