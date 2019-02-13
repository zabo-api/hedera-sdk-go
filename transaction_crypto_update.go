package hedera

// #include "hedera.h"
import "C"
import "time"

type TransactionCryptoUpdate struct {
	transaction
}

func newTransactionCryptoUpdate(client *Client, id AccountID) TransactionCryptoUpdate {
	return TransactionCryptoUpdate{transaction{
		C.hedera_transaction__crypto_update__new(client.inner, cAccountID(id))}}
}

func (tx TransactionCryptoUpdate) Key(public PublicKey) TransactionCryptoUpdate {
	C.hedera_transaction__crypto_update__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionCryptoUpdate) ProxyAccount(id AccountID) TransactionCryptoUpdate {
	C.hedera_transaction__crypto_update__set_proxy_account(tx.inner, cAccountID(id))
	return tx
}

func (tx TransactionCryptoUpdate) ProxyFraction(fraction int32) TransactionCryptoUpdate {
	C.hedera_transaction__crypto_update__set_proxy_fraction(tx.inner, C.int32_t(fraction))
	return tx
}

func (tx TransactionCryptoUpdate) SendRecordThreshold(threshold uint64) TransactionCryptoUpdate {
	C.hedera_transaction__crypto_update__set_send_record_threshold(tx.inner, C.uint64_t(threshold))
	return tx
}

func (tx TransactionCryptoUpdate) ReceiveRecordThreshold(threshold uint64) TransactionCryptoUpdate {
	C.hedera_transaction__crypto_update__set_receive_record_threshold(tx.inner, C.uint64_t(threshold))
	return tx
}

func (tx TransactionCryptoUpdate) AutoRenewPeriod(seconds uint64, nanos uint32) TransactionCryptoUpdate {

	cDuration := C.HederaDuration{
		seconds: C.uint64_t(seconds),
		nanos:   C.uint32_t(nanos),
	}
	C.hedera_transaction__crypto_update__set_auto_renew_period(tx.inner, cDuration)
	return tx
}

func (tx TransactionCryptoUpdate) ExpiresAt(expiresAt time.Time) TransactionCryptoUpdate {
	cExpiration := C.HederaTimestamp{
		seconds: C.int64_t(expiresAt.Unix()),
		nanos:   C.uint32_t(expiresAt.Nanosecond()),
	}

	C.hedera_transaction__crypto_update__set_expires_at(tx.inner, cExpiration)
	return tx
}

func (tx TransactionCryptoUpdate) Operator(id AccountID) TransactionCryptoUpdate {
	return TransactionCryptoUpdate{tx.transaction.Operator(id)}
}

func (tx TransactionCryptoUpdate) Node(id AccountID) TransactionCryptoUpdate {
	return TransactionCryptoUpdate{tx.transaction.Node(id)}
}

func (tx TransactionCryptoUpdate) Memo(memo string) TransactionCryptoUpdate {
	return TransactionCryptoUpdate{tx.transaction.Memo(memo)}
}
