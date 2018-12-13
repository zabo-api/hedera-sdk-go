#pragma once

#include "hedera-transaction.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__admin_file_recover__new(HederaClient*, HederaFileId);

extern HederaTransaction* hedera_transaction__admin_contract_recover__new(HederaClient*, HederaContractId);

#ifdef __cplusplus
}
#endif