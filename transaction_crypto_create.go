package hedera

// #include "hedera.h"
import "C"

type TransactionCryptoCreate struct {
	transaction
}

func newTransactionCryptoCreate(client *Client) TransactionCryptoCreate {
	return TransactionCryptoCreate{transaction{
		C.hedera_transaction__crypto_create__new(client.inner)}}
}

func (tx TransactionCryptoCreate) Key(public PublicKey) TransactionCryptoCreate {
	C.hedera_transaction__crypto_create__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionCryptoCreate) InitialBalance(balance uint64) TransactionCryptoCreate {
	C.hedera_transaction__crypto_create__set_initial_balance(tx.inner, C.uint64_t(balance))
	return tx
}

func (tx TransactionCryptoCreate) Operator(id AccountID) TransactionCryptoCreate {
	return TransactionCryptoCreate{tx.transaction.Operator(id)}
}

func (tx TransactionCryptoCreate) Node(id AccountID) TransactionCryptoCreate {
	return TransactionCryptoCreate{tx.transaction.Node(id)}
}

func (tx TransactionCryptoCreate) Memo(memo string) TransactionCryptoCreate {
	return TransactionCryptoCreate{tx.transaction.Memo(memo)}
}
