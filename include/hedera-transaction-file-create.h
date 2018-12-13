#pragma once

#include <stdint.h>
#include "hedera-transaction.h"
#include "hedera-timestamp.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaTransaction* hedera_transaction__file_create__new(HederaClient*);

extern void hedera_transaction__file_create__set_key(
    HederaTransaction*, HederaPublicKey key);

extern void hedera_transaction__file_create__set_contents(
    HederaTransaction*, const uint8_t* contents);

extern void hedera_transaction__file_create__set_expires_at(
    HederaClient*, HederaTimestamp expiration);

#ifdef __cplusplus
}
#endif
