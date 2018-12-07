#pragma once

#include "hedera-timestamp.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    HederaAccountId account_id;
    HederaTimestamp transaction_valid_start;
} HederaTransactionId;

extern HederaTransactionId hedera_transaction_id_new(HederaAccountId account);

extern char* hedera_transaction_id_to_str(HederaTransactionId*);

/// Parse a [HederaTransactionID] from a string.
extern HederaError hedera_transaction_id_from_str(const char* s, HederaTransactionId* out);

#ifdef __cplusplus
}
#endif
