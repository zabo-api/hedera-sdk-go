#pragma once

#include "hedera-transaction.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__file_delete__new(HederaClient*, HederaFileId id);

#ifdef __cplusplus
}
#endif