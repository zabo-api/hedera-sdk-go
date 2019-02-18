package hedera

// #include "hedera.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/markbates/oncer"
)

type QueryTransactionGetRecord struct {
	query
}

func newQueryTransactionGetRecord(client Client, id TransactionID) QueryTransactionGetRecord {
	return QueryTransactionGetRecord{
		query{C.hedera_query__transaction_get_record__new(client.inner, cTransactionID(id))},
	}
}

// Deprecated: Use Query.Get() instead
func (query QueryTransactionGetRecord) Answer() (TransactionRecord, error) {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Query.Answer()",
		"Use Query.Get() instead.")

	return query.Get()
}

func (query QueryTransactionGetRecord) Get() (TransactionRecord, error) {
	var out C.HederaTransactionRecord
	fmt.Printf("Query is = %v\n", query)
	res := C.hedera_query__transaction_get_record__get(query.inner, &out)

	if res != 0 {
		return TransactionRecord{}, hederaLastError()
	}

	record := TransactionRecord{}

	// if out.memo != nil {
	// 	p := out.memo
	// 	record.memo = C.GoString(p)
	// }

	// record.receipt = TransactionReceipt{Status: Status(out.receipt.status)}
	// if out.receipt.account_id != nil {
	// 	accountID := goAccountID(*out.receipt.account_id)
	// 	record.receipt.AccountID = &accountID
	// }

	// if out.receipt.contract_id != nil {
	// 	contractID := goContractID(*out.receipt.contract_id)
	// 	record.receipt.ContractID = &contractID
	// }

	// if out.receipt.file_id != nil {
	// 	fileID := goFileID(*out.receipt.file_id)
	// 	record.receipt.FileID = &fileID
	// }
	// length := uint64(out.transaction_hash.len)
	// record.transaction_hash = CArrayToByteSlice(out.transaction_hash.ptr, length)

	// record.transaction_fee = uint64(out.transaction_fee)

	// consensus_timestamp      time.Time
	// contract_function_result ContractFunctionResult
	// transfers                []Transfer

	return record, nil
}

func CArrayToByteSlice(array unsafe.Pointer, size uint64) []byte {
	var arrayptr = uintptr(array)
	var byteSlice = make([]byte, size)

	for i := 0; i < len(byteSlice); i++ {
		byteSlice[i] = byte(*(*C.uchar)(unsafe.Pointer(arrayptr)))
		arrayptr++
	}

	return byteSlice
}
