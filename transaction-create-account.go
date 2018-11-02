package hedera

// #include "hedera-transaction-create-account.h"
import "C"
import "unsafe"

type TransactionCreateAccount struct {
	Transaction
}

func newTransactionCreateAccount(client *Client) TransactionCreateAccount {
	return TransactionCreateAccount{Transaction{
		C.hedera_transaction__create_account__new(client.inner)}}
}

func (tx TransactionCreateAccount) Operator(id AccountID) TransactionCreateAccount {
	C.hedera_transaction_set_operator(tx.inner, id.c())
	return tx
}

func (tx TransactionCreateAccount) Node(id AccountID) TransactionCreateAccount {
	C.hedera_transaction_set_node(tx.inner, id.c())
	return tx
}

func (tx TransactionCreateAccount) Memo(memo string) TransactionCreateAccount {
	C.hedera_transaction_set_memo(tx.inner, C.CString(memo))
	return tx
}

func (tx TransactionCreateAccount) Sign(key SecretKey) TransactionCreateAccount {
	C.hedera_transaction_sign(tx.inner, key.inner)
	return tx
}

func (tx TransactionCreateAccount) Key(public PublicKey) TransactionCreateAccount {
	C.hedera_transaction__create_account__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionCreateAccount) InitialBalance(balance uint64) TransactionCreateAccount {
	C.hedera_transaction__create_account__set_initial_balance(tx.inner, C.uint64_t(balance))
	return tx
}

func (tx TransactionCreateAccount) Execute() (TransactionResponse, error) {
	var res C.HederaTransactionResponse
	err := C.hedera_transaction__create_account__execute(tx.inner, &res)

	return *((*TransactionResponse)(unsafe.Pointer(&res))), hederaError(err)
}
