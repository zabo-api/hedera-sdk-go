#pragma once

#include "hedera-transaction.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__contract_update__new(HederaClient*);

extern void hedera_transaction__contract_update__set_file(HederaTransaction*, HederaFileId id);

extern void hedera_transaction__contract_update__set_admin_key(
    HederaTransaction*, HederaPublicKey key);

extern void hedera_transaction__contract_update__set_proxy_account(
    HederaTransaction*, HederaAccountId id);

#ifdef __cplusplus
}
#endif
