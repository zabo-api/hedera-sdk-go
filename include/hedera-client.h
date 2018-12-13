#pragma  once

#include "hedera-id.h"
#include "hedera-error.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
    // response codes for pre check validation
    STATUS_OK = 0, // the transaction passed the precheck
    STATUS_INVALID_TRANSACTION = 1, // For any error not handled by specific error codes listed below.
    STATUS_PAYER_ACCOUNT_NOT_FOUND  = 2, //Payer account does not exist.
    STATUS_INVALID_NODE_ACCOUNT=3, //Node Account provided does not match the node account of the node the transaction was submitted to.
    STATUS_TRANSACTION_EXPIRED = 4, // Pre-Check TransactionValidStart + transactionValidDuration is less than current consensus time.
    STATUS_INVALID_TRANSACTION_START = 5,// Transaction start time is greater than current consensus time
    STATUS_INVALID_TRANSACTION_DURATION = 6,//valid transaction duration is a positive non zero number that does not exceed 120 seconds
    STATUS_INVALID_SIGNATURE = 7,//the transaction signature is not valid
    STATUS_MEMO_TOO_LONG = 8,//Transaction memo size exceeded 100 bytes
    STATUS_INSUFFICIENT_TX_FEE  = 9, // the transaction fee is insufficient for this type of transaction
    STATUS_INSUFFICIENT_PAYER_BALANCE  = 10, // the payer account has insufficient cryptocurrency to pay the transaction fee
    STATUS_DUPLICATE_TRANSACTION = 11, // this transaction ID is a duplicate of one that was submitted to this node or reached consensus in the last 180 seconds (receipt period)
    STATUS_BUSY = 12, //If API is throttled out
    STATUS_NOT_SUPPORTED = 13, //not supported API

    //response codes used in queries
    STATUS_INVALID_FILE_ID = 14,//the file id is invalid or does not exist
    STATUS_INVALID_ACCOUNT_ID = 15,//the account id is invalid or does not exist
    STATUS_INVALID_CONTRACT_ID = 16,//the contract id is invalid or does ont exist
    STATUS_INVALID_TRANSACTION_ID =17,//transaction id is not valid
    STATUS_RECEIPT_NOT_FOUND = 18,//receipt for given transaction id does not exist
    STATUS_RECORD_NOT_FOUND = 19,//record for given transaction id does not exist
    STATUS_INVALID_SOLIDITY_ID = 20,//the solidity id is invalid or entity with this solidity id does not exist

    // response code for Transaction receipt
    STATUS_UNKNOWN = 21, // hasn't yet reached consensus, or has already expired
    STATUS_SUCCESS = 22, // the transaction succeeded
    STATUS_FAIL_INVALID = 23, // the transaction failed because it is invalid
    STATUS_FAIL_FEE = 24, // the transaction fee was insufficient
    STATUS_FAIL_BALANCE = 25, // the paying account had insufficient cryptocurrency

    //Crypto Response codes
    STATUS_KEY_REQUIRED = 26, //Key not provided in the transaction body
    STATUS_BAD_ENCODING = 27, //Unsupported algorithm/encoding used for keys in the transaction
    STATUS_INSUFFICIENT_ACCOUNT_BALANCE = 28, //When the account balance is not sufficient for the transfer
    STATUS_INVALID_SOLIDITY_ADDRESS = 29, //During an update transaction when the system is not able to find the Users Solidity address

    //Smart contract creation or execution
    STATUS_INSUFFICIENT_GAS = 30, //Not enough gas was supplied to execute tranasction
    STATUS_CONTRACT_SIZE_LIMIT_EXCEEDED =31, //contract byte code size is over the limit
    STATUS_LOCAL_CALL_MODIFICATION_EXCEPTION=32,//local execution (query) is requested for a function which changes state
    STATUS_CONTRACT_REVERT_EXECUTED=33, //Contract REVERT OPCODE executed
    STATUS_CONTRACT_EXECUTION_EXCEPTION=34,//For any contract execution related error not handled by specific error codes listed above.
    STATUS_INVALID_RECEIVING_NODE_ACCOUNT = 35, //In Query validation, account with +ve(amount) value should be Receiving node account, the receiver account should be only one account in the list
    STATUS_MISSING_QUERY_HEADER = 36, // Header is missing in Query request
    STATUS_ACCOUNT_UPDATE_FAILED = 37, // the update of the account failed
    STATUS_INVALID_KEY_ENCODING = 38,
    STATUS_NULL_SOLIDITY_ADDRESS = 39, // null solidity address
    STATUS_CONTRACT_UPDATE_FAILED = 40, // update of the contract failed
    STATUS_INVALID_QUERY_HEADER = 41, // the query header is invalid
    STATUS_INVALID_FEE_SUBMITTED = 42, // Invalid fee submitted*/
    STATUS_INVALID_PAYER_SIGNATURE = 43, //  payer signature is invalid

    STATUS_KEY_NOT_PROVIDED = 44,
    STATUS_INVALID_EXPIRATION_TIME = 45,
    STATUS_NO_WACL_KEY = 46,
    STATUS_FILE_CONTENT_EMPTY = 47,
    STATUS_INVALID_ACCOUNT_AMOUNTS = 48, // The crypto transfer credit and debit don't equal to 0

    // new response codes to be added
    STATUS_EMPTY_TRANSACTION_BODY = 49, // transaction body is empty
    STATUS_INVALID_TRANSACTION_BODY = 50, // invalid transaction body
} HederaStatus;

typedef struct HederaClient HederaClient;

/// Establish a connection to a Hedera node.
/// Must be closed with [hedera_client_close].
extern HederaError hedera_client_new(const char* address, HederaClient**);

/// Close and releases resources for a [HederaClient].
extern void hedera_client_close(HederaClient*);

#ifdef __cplusplus
}
#endif
