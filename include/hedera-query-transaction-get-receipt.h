#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-transaction-id.h"
#include "hedera-query.h"
#include "hedera-transaction-receipt.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__transaction_get_receipt__new(
    HederaClient*,
    HederaTransactionId transaction_id
);

extern HederaError hedera_query__transaction_get_receipt__get(
    HederaQuery*, 
    HederaTransactionReceipt**);

#ifdef __cplusplus
}
#endif
