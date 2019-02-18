package hedera

type TransactionRecord struct {
	// receipt                  TransactionReceipt
	transaction_hash []byte
	// consensus_timestamp      time.Time
	// memo                     string
	// transaction_fee          uint64
	// contract_function_result ContractFunctionResult
	// transfers                []Transfer
}

type ContractFunctionResult struct {
	ContractID           *ContractID
	contract_call_result []byte
	error_message        string
	bloom                []byte
	gas_used             uint64
	log_info             []byte
}

type ContractLogInfo struct {
	bloom []byte
	topic []byte
	data  []byte
}

type Transfer struct {
	AccountID *AccountID
	amount    int64
}
