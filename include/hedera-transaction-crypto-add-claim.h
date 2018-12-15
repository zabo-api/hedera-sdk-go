#pragma once

#include "hedera-transaction.h"
#include "hedera-claim.h"
#include "hedera-id.h"
#include "hedera-array.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__crypto_add_claim__new(
    HederaClient*,
    HederaAccountId account_id,
    HederaArray hash
);

extern void hedera_transaction__crypto_add_claim__add_key(
    HederaTransaction*, HederaPublicKey key);

#ifdef __cplusplus
}
#endif
