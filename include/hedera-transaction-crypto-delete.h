#pragma once

#include "hedera-transaction.h"
#include "hedera-claim.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__crypto_delete__new(
    HederaClient*, 
    HederaAccountId account_id
);

extern void hedera_transaction__crypto_delete__set_transfer_to(
    HederaTransaction*, 
    HederaAccountId transfer_to
);

#ifdef __cplusplus
}
#endif
