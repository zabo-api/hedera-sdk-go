#pragma once

#include "hedera-transaction.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__crypto_create__new(HederaClient*);

extern void hedera_transaction__crypto_create__set_initial_balance(HederaTransaction*, uint64_t balance);

extern void hedera_transaction__crypto_create__set_key(HederaTransaction*, HederaPublicKey key);

#ifdef __cplusplus
}
#endif
