package hedera

type TransactionResponse uint8

const (
	TransactionResponseOk TransactionResponse = 0
	TransactionResponseInvalidTransaction TransactionResponse = 1
	TransactionResponsePayerAccountNotFound TransactionResponse = 2
	TransactionResponseInvalidNodeAccount TransactionResponse = 3
	TransactionResponseTransactionExpired TransactionResponse = 4
	TransactionResponseInvalidTransactionStart TransactionResponse = 5
	TransactionResponseInvalidTransactionDuration TransactionResponse = 6
	TransactionResponseInvalidSignature TransactionResponse = 7
	TransactionResponseMemoTooLong TransactionResponse = 8
	TransactionResponseInsufficientTxFee TransactionResponse = 9
	TransactionResponseInsufficientPayerBalance TransactionResponse = 10
	TransactionResponseDuplicateTransaction TransactionResponse = 11
	TransactionResponseBusy TransactionResponse = 12
	TransactionResponseNotSupported TransactionResponse = 13
	TransactionResponseInvalidFileId TransactionResponse = 14
	TransactionResponseInvalidAccountId TransactionResponse = 15
	TransactionResponseInvalidContractId TransactionResponse = 16
	TransactionResponseInvalidTransactionId TransactionResponse = 17
	TransactionResponseReceiptNotFound TransactionResponse = 18
	TransactionResponseRecordNotFound TransactionResponse = 19
	TransactionResponseInvalidSolidityId TransactionResponse = 20
	TransactionResponseUnknown TransactionResponse = 21
	TransactionResponseSuccess TransactionResponse = 22
	TransactionResponseFailInvalid TransactionResponse = 23
	TransactionResponseFailFee TransactionResponse = 24
	TransactionResponseFailBalance TransactionResponse = 25
	TransactionResponseKeyRequired TransactionResponse = 26
	TransactionResponseBadEncoding TransactionResponse = 27
	TransactionResponseInsufficientAccountBalance TransactionResponse = 28
	TransactionResponseInvalidSolidityAddress TransactionResponse = 29
	TransactionResponseInsufficientGas TransactionResponse = 30
	TransactionResponseContractSizeLimitExceeded TransactionResponse = 31
	TransactionResponseLocalCallModificationException TransactionResponse = 32
	TransactionResponseContractRevertExecuted TransactionResponse = 33
	TransactionResponseContractExecutionException TransactionResponse = 34
	TransactionResponseInvalidReceivingNodeAccount TransactionResponse = 35
	TransactionResponseMissingQueryHeader TransactionResponse = 36
	TransactionResponseAccountUpdateFailed TransactionResponse = 37
	TransactionResponseInvalidKeyEncoding TransactionResponse = 38
	TransactionResponseNullSolidityAddress TransactionResponse = 39
	TransactionResponseContractUpdateFailed TransactionResponse = 40
	TransactionResponseInvalidQueryHeader TransactionResponse = 41
	TransactionResponseInvalidFeeSubmitted TransactionResponse = 42
	TransactionResponseInvalidPayerSignature TransactionResponse = 43
	TransactionResponseKeyNotProvided TransactionResponse = 44
	TransactionResponseInvalidExpirationTime TransactionResponse = 45
	TransactionResponseNoWaclKey TransactionResponse = 46
	TransactionResponseFileContentEmpty TransactionResponse = 47
	TransactionResponseInvalidAccountAmounts TransactionResponse = 48
	TransactionResponseEmptyTransactionBody TransactionResponse = 49
	TransactionResponseInvalidTransactionBody TransactionResponse = 50
)

var transactionResponseText = map[TransactionResponse]string{
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
