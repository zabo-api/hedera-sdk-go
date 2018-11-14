package hedera

// #include "hedera-transaction-crypto-transfer.h"
import "C"

type TransactionCryptoTransfer struct {
	Transaction
}

func newTransactionCryptoTransfer(client *Client) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{Transaction{
		C.hedera_transaction__crypto_transfer__new(client.inner)}}
}

func (tx TransactionCryptoTransfer) Transfer(id AccountID, amount int64) TransactionCryptoTransfer {
	C.hedera_transaction__crypto_transfer__add_transfer(tx.inner, id.c(), C.int64_t(amount))
	return tx
}

//
// Inherited from Transaction
//

func (tx TransactionCryptoTransfer) Operator(id AccountID) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.Transaction.Operator(id)}
}

func (tx TransactionCryptoTransfer) Node(id AccountID) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.Transaction.Node(id)}
}

func (tx TransactionCryptoTransfer) Memo(memo string) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.Transaction.Memo(memo)}
}

func (tx TransactionCryptoTransfer) Sign(key SecretKey) TransactionCryptoTransfer {
	return TransactionCryptoTransfer{tx.Transaction.Sign(key)}
}
