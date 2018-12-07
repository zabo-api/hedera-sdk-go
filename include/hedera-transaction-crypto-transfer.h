#pragma once

#include "hedera-transaction.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__crypto_transfer__new(HederaClient*);

extern void hedera_transaction__crypto_transfer__add_transfer(HederaTransaction*, HederaAccountId id, int64_t amount);

#ifdef __cplusplus
}
#endif
