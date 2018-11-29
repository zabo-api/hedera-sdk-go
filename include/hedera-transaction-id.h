#ifndef HEDERA_TRANSACTION_ID_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
#define HEDERA_TRANSACTION_ID_9999A0E8_2BD1_4C33_8071_D93A13B8A9E

#include "hedera-timestamp.h"
#include "hedera-account-id.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    HederaAccountId account_id;
    HederaTimestamp transaction_valid_start;
} HederaTransactionId;

extern HederaTransactionId hedera_transaction_id_new(HederaAccountId account);

extern char* hedera_transaction_id_to_str(HederaTransactionId);

/// Parse a [HederaTransactionID] from a string.
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_transaction_id_from_str(const char* s, HederaTransactionId* out);

#ifdef __cplusplus
}
#endif

#endif // HEDERA_TRANSACTION_ID_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
