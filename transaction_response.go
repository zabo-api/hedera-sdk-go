package hedera

type TransactionResponse uint8

const (
	TransactionResponseOK TransactionResponse = 0
	TransactionResponseINVALIDTRANSACTION TransactionResponse = 1
	TransactionResponsePAYERACCOUNTNOTFOUND  TransactionResponse = 2
	TransactionResponseINVALIDNODEACCOUNT TransactionResponse = 3
	TransactionResponseTRANSACTIONEXPIRED TransactionResponse = 4
	TransactionResponseINVALIDTRANSACTIONSTART TransactionResponse = 5
	TransactionResponseINVALIDTRANSACTIONDURATION TransactionResponse = 6
	TransactionResponseINVALIDSIGNATURE TransactionResponse = 7
	TransactionResponseMEMOTOOLONG TransactionResponse = 8
	TransactionResponseINSUFFICIENTTXFEE  TransactionResponse = 9
	TransactionResponseINSUFFICIENTPAYERBALANCE  TransactionResponse = 10
	TransactionResponseDUPLICATETRANSACTION TransactionResponse = 11
	TransactionResponseBUSY TransactionResponse = 12
	TransactionResponseNOTSUPPORTED TransactionResponse = 13

	TransactionResponseINVALIDFILEID TransactionResponse = 14
	TransactionResponseINVALIDACCOUNTID TransactionResponse = 15
	TransactionResponseINVALIDCONTRACTID TransactionResponse = 16
	TransactionResponseINVALIDTRANSACTIONID TransactionResponse = 17
	TransactionResponseRECEIPTNOTFOUND TransactionResponse = 18
	TransactionResponseRECORDNOTFOUND TransactionResponse = 19
	TransactionResponseINVALIDSOLIDITYID TransactionResponse = 20

	TransactionResponseUNKNOWN TransactionResponse = 21
	TransactionResponseSUCCESS TransactionResponse = 22
	TransactionResponseFAILINVALID TransactionResponse = 23
	TransactionResponseFAILFEE TransactionResponse = 24
	TransactionResponseFAILBALANCE TransactionResponse = 25

	TransactionResponseKEYREQUIRED TransactionResponse = 26
	TransactionResponseBADENCODING TransactionResponse = 27
	TransactionResponseINSUFFICIENTACCOUNTBALANCE TransactionResponse = 28
	TransactionResponseINVALIDSOLIDITYADDRESS TransactionResponse = 29

	TransactionResponseINSUFFICIENTGAS TransactionResponse = 30
	TransactionResponseCONTRACTSIZELIMITEXCEEDED TransactionResponse = 31
	TransactionResponseLOCALCALLMODIFICATIONEXCEPTION TransactionResponse = 32
	TransactionResponseCONTRACTREVERTEXECUTED TransactionResponse = 33
	TransactionResponseCONTRACTEXECUTIONEXCEPTION TransactionResponse= 34
	TransactionResponseINVALIDRECEIVINGNODEACCOUNT TransactionResponse = 35
	TransactionResponseMISSINGQUERYHEADER TransactionResponse = 36
	TransactionResponseACCOUNTUPDATEFAILED TransactionResponse = 37
	TransactionResponseINVALIDKEYENCODING TransactionResponse = 38
	TransactionResponseNULLSOLIDITYADDRESS TransactionResponse = 39
	TransactionResponseCONTRACTUPDATEFAILED TransactionResponse = 40
	TransactionResponseINVALIDQUERYHEADER TransactionResponse = 41
	TransactionResponseINVALIDFEESUBMITTED TransactionResponse = 42
	TransactionResponseINVALIDPAYERSIGNATURE TransactionResponse = 43

	TransactionResponseKEYNOTPROVIDED TransactionResponse = 44
	TransactionResponseINVALIDEXPIRATIONTIME TransactionResponse = 45
	TransactionResponseNOWACLKEY TransactionResponse = 46
	TransactionResponseFILECONTENTEMPTY TransactionResponse = 47
	TransactionResponseINVALIDACCOUNTAMOUNTS TransactionResponse = 48

	TransactionResponseEMPTYTRANSACTIONBODY TransactionResponse = 49
	TransactionResponseINVALIDTRANSACTIONBODY TransactionResponse = 50
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
