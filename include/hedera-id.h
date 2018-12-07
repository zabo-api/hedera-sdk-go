#pragma once

#include <stdint.h>
#include "hedera-error.h"

#ifdef __cplusplus
extern "C" {
#endif

/// The ID for a crypto-currency account.
typedef struct {
    int64_t shard;      ///< The shard number (non-negative).
    int64_t realm;      ///< The realm number (non-negative).
    int64_t account;    ///< A non-negative number unique within its realm.
} HederaAccountId;

extern HederaError hedera_account_id_from_str(const char* s, HederaAccountId* out);

extern const char* hedera_account_id_to_str(HederaAccountId*);

/// The ID for a smart contract instance.
typedef struct {
    int64_t shard;      ///< The shard number (non-negative).
    int64_t realm;      ///< The realm number (non-negative).
    int64_t contract;   ///< A non-negative number unique within its realm.
} HederaContractId;

extern HederaError hedera_contract_id_from_str(const char* s, HederaContractId* out);

extern const char* hedera_contract_id_to_str(HederaContractId*);

/// The ID for a file.
typedef struct {
    int64_t shard;      ///< The shard number (non-negative).
    int64_t realm;      ///< The realm number (non-negative).
    int64_t file;       ///< A non-negative number unique within its realm.
} HederaFileId;

extern HederaError hedera_file_id_from_str(const char* s, HederaFileId* out);

extern const char* hedera_file_id_to_str(HederaFileId*);

#ifdef __cplusplus
}
#endif
