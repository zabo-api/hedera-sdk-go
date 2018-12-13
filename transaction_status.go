package hedera

type TransactionStatus uint8

const (
	TransactionStatusOk                             TransactionStatus = 0
	TransactionStatusInvalidTransaction             TransactionStatus = 1
	TransactionStatusPayerAccountNotFound           TransactionStatus = 2
	TransactionStatusInvalidNodeAccount             TransactionStatus = 3
	TransactionStatusTransactionExpired             TransactionStatus = 4
	TransactionStatusInvalidTransactionStart        TransactionStatus = 5
	TransactionStatusInvalidTransactionDuration     TransactionStatus = 6
	TransactionStatusInvalidSignature               TransactionStatus = 7
	TransactionStatusMemoTooLong                    TransactionStatus = 8
	TransactionStatusInsufficientTxFee              TransactionStatus = 9
	TransactionStatusInsufficientPayerBalance       TransactionStatus = 10
	TransactionStatusDuplicateTransaction           TransactionStatus = 11
	TransactionStatusBusy                           TransactionStatus = 12
	TransactionStatusNotSupported                   TransactionStatus = 13
	TransactionStatusInvalidFileId                  TransactionStatus = 14
	TransactionStatusInvalidAccountId               TransactionStatus = 15
	TransactionStatusInvalidContractId              TransactionStatus = 16
	TransactionStatusInvalidTransactionId           TransactionStatus = 17
	TransactionStatusReceiptNotFound                TransactionStatus = 18
	TransactionStatusRecordNotFound                 TransactionStatus = 19
	TransactionStatusInvalidSolidityId              TransactionStatus = 20
	TransactionStatusUnknown                        TransactionStatus = 21
	TransactionStatusSuccess                        TransactionStatus = 22
	TransactionStatusFailInvalid                    TransactionStatus = 23
	TransactionStatusFailFee                        TransactionStatus = 24
	TransactionStatusFailBalance                    TransactionStatus = 25
	TransactionStatusKeyRequired                    TransactionStatus = 26
	TransactionStatusBadEncoding                    TransactionStatus = 27
	TransactionStatusInsufficientAccountBalance     TransactionStatus = 28
	TransactionStatusInvalidSolidityAddress         TransactionStatus = 29
	TransactionStatusInsufficientGas                TransactionStatus = 30
	TransactionStatusContractSizeLimitExceeded      TransactionStatus = 31
	TransactionStatusLocalCallModificationException TransactionStatus = 32
	TransactionStatusContractRevertExecuted         TransactionStatus = 33
	TransactionStatusContractExecutionException     TransactionStatus = 34
	TransactionStatusInvalidReceivingNodeAccount    TransactionStatus = 35
	TransactionStatusMissingQueryHeader             TransactionStatus = 36
	TransactionStatusAccountUpdateFailed            TransactionStatus = 37
	TransactionStatusInvalidKeyEncoding             TransactionStatus = 38
	TransactionStatusNullSolidityAddress            TransactionStatus = 39
	TransactionStatusContractUpdateFailed           TransactionStatus = 40
	TransactionStatusInvalidQueryHeader             TransactionStatus = 41
	TransactionStatusInvalidFeeSubmitted            TransactionStatus = 42
	TransactionStatusInvalidPayerSignature          TransactionStatus = 43
	TransactionStatusKeyNotProvided                 TransactionStatus = 44
	TransactionStatusInvalidExpirationTime          TransactionStatus = 45
	TransactionStatusNoWaclKey                      TransactionStatus = 46
	TransactionStatusFileContentEmpty               TransactionStatus = 47
	TransactionStatusInvalidAccountAmounts          TransactionStatus = 48
	TransactionStatusEmptyTransactionBody           TransactionStatus = 49
	TransactionStatusInvalidTransactionBody         TransactionStatus = 50
)

var transactionStatusText = map[TransactionStatus]string{
	0 : "OK",
	1 : "INVALIDTRANSACTION",
	2 : "PAYERACCOUNTNOTFOUND",
	3 : "INVALIDNODEACCOUNT",
	4 : "TRANSACTIONEXPIRED",
	5 : "INVALIDTRANSACTIONSTART",
	6 : "INVALIDTRANSACTIONDURATION",
	7 : "INVALIDSIGNATURE",
	8 : "MEMOTOOLONG",
	9 : "INSUFFICIENTTXFEE",
	10: "INSUFFICIENTPAYERBALANCE",
	11: "DUPLICATETRANSACTION",
	12: "BUSY",
	13: "NOTSUPPORTED",
	14: "INVALIDFILEID",
	15: "INVALIDACCOUNTID",
	16: "INVALIDCONTRACTID",
	17: "INVALIDTRANSACTIONID",
	18: "RECEIPTNOTFOUND",
	19: "RECORDNOTFOUND",
	20: "INVALIDSOLIDITYID",
	21: "UNKNOWN",
	22: "SUCCESS",
	23: "FAILINVALID",
	24: "FAILFEE",
	25: "FAILBALANCE",
	26: "KEYREQUIRED",
	27: "BADENCODING",
	28: "INSUFFICIENTACCOUNTBALANCE",
	29: "INVALIDSOLIDITYADDRESS",
	30: "INSUFFICIENTGAS",
	31: "CONTRACTSIZELIMITEXCEEDED",
	32: "LOCALCALLMODIFICATIONEXCEPTION",
	33: "CONTRACTREVERTEXECUTED",
	34: "CONTRACTEXECUTIONEXCEPTION",
	35: "INVALIDRECEIVINGNODEACCOUNT",
	36: "MISSINGQUERYHEADER",
	37: "ACCOUNTUPDATEFAILED",
	38: "INVALIDKEYENCODING",
	39: "NULLSOLIDITYADDRESS",
	40: "CONTRACTUPDATEFAILED",
	41: "INVALIDQUERYHEADER",
	42: "INVALIDFEESUBMITTED",
	43: "INVALIDPAYERSIGNATURE",
	44: "KEYNOTPROVIDED",
	45: "INVALIDEXPIRATIONTIME",
	46: "NOWACLKEY",
	47: "FILECONTENTEMPTY",
	48: "INVALIDACCOUNTAMOUNTS",
	49: "EMPTYTRANSACTIONBODY",
	50: "INVALIDTRANSACTIONBODY",
}
