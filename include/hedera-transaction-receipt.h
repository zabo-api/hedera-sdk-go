#pragma once

#include <stdint.h>
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    // fixme: use enum here
    uint8_t status;
    HederaAccountId* account_id;
    HederaContractId* contract_id;
    HederaFileId* file_id;
} HederaTransactionReceipt;

#ifdef __cplusplus
}
#endif
