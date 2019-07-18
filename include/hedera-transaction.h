#pragma once

#include <stdint.h>

#include "hedera-error.h"
#include "hedera-crypto.h"
#include "hedera-client.h"
#include "hedera-id.h"
#include "hedera-transaction-id.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct HederaTransaction HederaTransaction;

extern void hedera_transaction_set_operator(HederaTransaction*, HederaAccountId operator_);

extern void hedera_transaction_set_node(HederaTransaction*, HederaAccountId node);

extern void hedera_transaction_set_memo(HederaTransaction*, const char* memo);

extern void hedera_transaction_set_fee(HederaTransaction*, uint64_t fee);

extern void hedera_transaction_sign(HederaTransaction*, HederaSecretKey*);

extern HederaError hedera_transaction_execute(HederaTransaction*, HederaTransactionId*);

#ifdef __cplusplus
}
#endif
