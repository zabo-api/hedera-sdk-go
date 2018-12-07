#pragma  once

#include "hedera-id.h"
#include "hedera-error.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
    /// The transaction passed the precheck.
    PRECHECK_OK = 0,

    /// The transaction had incorrect syntax or other errors.
    PRECHECK_INVALID_TRANSACTION = 1,

    /// The payer account or node account isn't a valid account number.
    PRECHECK_INVALID_ACCOUNT = 2,

    /// The transaction fee is insufficient for this type of transaction.
    PRECHECK_INSUFFICIENT_FEE = 3,

    /// The payer account has an insufficient balance to pay the transaction fee.
    PRECHECK_INSUFFICIENT_BALANCE = 4,

    /// This transaction ID is a duplicate of one that was submitted to this node or
    /// reached consensus in the last 180 seconds (receipt period).
    PRECHECK_DUPLICATE = 5,

    /// If API is throttled out.
    PRECHECK_BUSY = 6,

    /// Unsupported request.
    PRECHECK_NOT_SUPPORTED = 7,
} HederaPrecheckCode;

typedef struct HederaClient HederaClient;

/// Establish a connection to a Hedera node.
/// Must be closed with [hedera_client_close].
extern HederaError hedera_client_new(const char* address, HederaClient**);

/// Close and releases resources for a [HederaClient].
extern void hedera_client_close(HederaClient*);

#ifdef __cplusplus
}
#endif
