package hedera

type Status uint8

const (
	StatusOk                             Status = 0
	StatusInvalidTransaction             Status = 1
	StatusPayerAccountNotFound           Status = 2
	StatusInvalidNodeAccount             Status = 3
	StatusTransactionExpired             Status = 4
	StatusInvalidTransactionStart        Status = 5
	StatusInvalidTransactionDuration     Status = 6
	StatusInvalidSignature               Status = 7
	StatusMemoTooLong                    Status = 8
	StatusInsufficientTxFee              Status = 9
	StatusInsufficientPayerBalance       Status = 10
	StatusDuplicateTransaction           Status = 11
	StatusBusy                           Status = 12
	StatusNotSupported                   Status = 13
	StatusInvalidFileId                  Status = 14
	StatusInvalidAccountId               Status = 15
	StatusInvalidContractId              Status = 16
	StatusInvalidTransactionId           Status = 17
	StatusReceiptNotFound                Status = 18
	StatusRecordNotFound                 Status = 19
	StatusInvalidSolidityId              Status = 20
	StatusUnknown                        Status = 21
	StatusSuccess                        Status = 22
	StatusFailInvalid                    Status = 23
	StatusFailFee                        Status = 24
	StatusFailBalance                    Status = 25
	StatusKeyRequired                    Status = 26
	StatusBadEncoding                    Status = 27
	StatusInsufficientAccountBalance     Status = 28
	StatusInvalidSolidityAddress         Status = 29
	StatusInsufficientGas                Status = 30
	StatusContractSizeLimitExceeded      Status = 31
	StatusLocalCallModificationException Status = 32
	StatusContractRevertExecuted         Status = 33
	StatusContractExecutionException     Status = 34
	StatusInvalidReceivingNodeAccount    Status = 35
	StatusMissingQueryHeader             Status = 36
	StatusAccountUpdateFailed            Status = 37
	StatusInvalidKeyEncoding             Status = 38
	StatusNullSolidityAddress            Status = 39
	StatusContractUpdateFailed           Status = 40
	StatusInvalidQueryHeader             Status = 41
	StatusInvalidFeeSubmitted            Status = 42
	StatusInvalidPayerSignature          Status = 43
	StatusKeyNotProvided                 Status = 44
	StatusInvalidExpirationTime          Status = 45
	StatusNoWaclKey                      Status = 46
	StatusFileContentEmpty               Status = 47
	StatusInvalidAccountAmounts          Status = 48
	StatusEmptyTransactionBody           Status = 49
	StatusInvalidTransactionBody         Status = 50
)

var statusText = map[Status]string{
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
