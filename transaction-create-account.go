package hedera

// #include "hedera-transaction-create-account.h"
import "C"

type TransactionCreateAccount struct {
	Transaction
}

func newTransactionCreateAccount(client *Client) TransactionCreateAccount {
	return TransactionCreateAccount{Transaction{
		C.hedera_transaction__create_account__new(client.inner)}}
}

func (tx TransactionCreateAccount) Key(public PublicKey) TransactionCreateAccount {
	C.hedera_transaction__create_account__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionCreateAccount) InitialBalance(balance uint64) TransactionCreateAccount {
	C.hedera_transaction__create_account__set_initial_balance(tx.inner, C.uint64_t(balance))
	return tx
}

//
// Inherited from Transaction
//

func (tx TransactionCreateAccount) Operator(id AccountID) TransactionCreateAccount {
	return TransactionCreateAccount{tx.Transaction.Operator(id)}
}

func (tx TransactionCreateAccount) Node(id AccountID) TransactionCreateAccount {
	return TransactionCreateAccount{tx.Transaction.Node(id)}
}

func (tx TransactionCreateAccount) Memo(memo string) TransactionCreateAccount {
	return TransactionCreateAccount{tx.Transaction.Memo(memo)}
}

func (tx TransactionCreateAccount) Sign(key SecretKey) TransactionCreateAccount {
	return TransactionCreateAccount{tx.Transaction.Sign(key)}
}
