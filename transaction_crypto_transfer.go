package hedera

// #include "hedera.h"
import "C"

type TransactionCryptoTransfer struct {
	transaction
}

func newTransactionCryptoTransfer(client *Client) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{transaction{
		C.hedera_transaction__crypto_transfer__new(client.inner)}}
}

func (tx TransactionCryptoTransfer) Transfer(id AccountID, amount int64) TransactionCryptoTransfer {
	C.hedera_transaction__crypto_transfer__add_transfer(tx.inner, cAccountID(id), C.int64_t(amount))
	return tx
}

func (tx TransactionCryptoTransfer) Operator(id AccountID) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.transaction.Operator(id)}
}

func (tx TransactionCryptoTransfer) Node(id AccountID) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.transaction.Node(id)}
}

func (tx TransactionCryptoTransfer) Memo(memo string) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.transaction.Memo(memo)}
}
