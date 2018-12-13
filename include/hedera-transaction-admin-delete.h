#pragma once

#include "hedera-transaction.h"
#include "hedera-id.h"
#include "hedera-timestamp.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__admin_file_delete__new(HederaClient*, HederaFileId);

extern void hedera_transaction__admin_file_delete__set_expire_at(HederaClient*, HederaTimestamp time);

extern HederaTransaction* hedera_transaction__admin_contract_delete__new(HederaClient*, HederaContractId);

extern void hedera_transaction__admin_contract_delete__set_expire_at(HederaClient*, HederaTimestamp time);

#ifdef __cplusplus
}
#endif
