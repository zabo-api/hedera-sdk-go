#pragma once

#include <stdint.h>
#include "hedera-transaction.h"
#include "hedera-id.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__file_append__new(
    HederaClient*, HederaFileId id, const uint8_t* contents);

#ifdef __cplusplus
}
#endif
