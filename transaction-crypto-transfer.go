package hedera

// #include "hedera-transaction-crypto-transfer.h"
import "C"
import "unsafe"

type TransactionCryptoTransfer struct {
	Transaction
}

func newTransactionCryptoTransfer(client *Client) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{Transaction{
		C.hedera_transaction__crypto_transfer__new(client.inner)}}
}

func (tx TransactionCryptoTransfer) Operator(id AccountID) TransactionCryptoTransfer {
	C.hedera_transaction_set_operator(tx.inner, id.c())
	return tx
}

func (tx TransactionCryptoTransfer) Node(id AccountID) TransactionCryptoTransfer {
	C.hedera_transaction_set_node(tx.inner, id.c())
	return tx
}

func (tx TransactionCryptoTransfer) Memo(memo string) TransactionCryptoTransfer {
	C.hedera_transaction_set_memo(tx.inner, C.CString(memo))
	return tx
}

func (tx TransactionCryptoTransfer) Sign(key SecretKey) TransactionCryptoTransfer {
	C.hedera_transaction_sign(tx.inner, key.inner)
	return tx
}

func (tx TransactionCryptoTransfer) Transfer(id AccountID, amount int64) TransactionCryptoTransfer {
	C.hedera_transaction__crypto_transfer__add_transfer(tx.inner, id.c(), C.int64_t(amount))
	return tx
}

func (tx TransactionCryptoTransfer) Execute() (TransactionResponse, error) {
	var res C.HederaTransactionResponse
	err := C.hedera_transaction__crypto_transfer__execute(tx.inner, &res)

	return *((*TransactionResponse)(unsafe.Pointer(&res))), hederaError(err)
}
