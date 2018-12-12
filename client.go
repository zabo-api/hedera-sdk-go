package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import (
	"unsafe"

	"github.com/markbates/oncer"
)

type Client struct {
	inner *C.HederaClient
}

func Dial(address string) (Client, error) {
	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))

	var inner *C.HederaClient
	res := C.hedera_client_new(cAddress, &inner)
	if res != 0 {
		return Client{}, hederaLastError()
	}

	return Client{inner}, nil
}

func (client Client) Close() {
	C.hedera_client_close(client.inner)
}

// Deprecated: Use Client.TransferCrypto() instead
func (client Client) CryptoTransfer() TransactionCryptoTransfer {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Client.CryptoTransfer()",
		"Use Client.TransferCrypto() instead.")

	return client.TransferCrypto()
}

// Deprecated: Use Client.Account(id).Balance() instead
func (client Client) GetAccountBalance(id AccountID) QueryCryptoGetAccountBalance {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Client.GetAccountBalance(id)",
		"Use Client.Account(id).Balance() instead.")

	return newQueryCryptoGetAccountBalance(client, id)
}

// Deprecated: Use Client.Transaction(id).Receipt() instead
func (client Client) GetTransactionReceipt(id *TransactionID) QueryGetTransactionReceipt {
	oncer.Deprecate(0,
		"github.com/hashgraph/hedera-sdk-go#Client.GetTransactionReceipt(id)",
		"Use Client.Transaction(id).Receipt() instead.")

	return newQueryGetTransactionReceipt(client, *id)
}

func (client Client) TransferCrypto() TransactionCryptoTransfer {
	return newTransactionCryptoTransfer(&client)
}

func (client Client) CreateAccount() TransactionCryptoCreate {
	return newTransactionCryptoCreate(&client)
}

func (client Client) Account(id AccountID) PartialAccountMessage {
	return PartialAccountMessage{client, id}
}

func (client Client) Transaction(id TransactionID) PartialTransactionMessage {
	return PartialTransactionMessage{client, id}
}

type PartialAccountMessage struct {
	client  Client
	account AccountID
}

func (m PartialAccountMessage) Balance() QueryCryptoGetAccountBalance {
	return newQueryCryptoGetAccountBalance(m.client, m.account)
}

type PartialTransactionMessage struct {
	client      Client
	transaction TransactionID
}

func (m PartialTransactionMessage) Receipt() QueryTransactionGetReceipt {
	return newQueryTransactionGetReceipt(m.client, m.transaction)
}
