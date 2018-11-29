package hedera

import "C"

// #include <stdlib.h>
// #include "hedera-key.h"
import "C"
import (
	"unsafe"
)

// An EdDSA secret key.
type SecretKey struct {
	inner C.HederaSecretKey
}

// A Signature signed by a Secret Key
type Signature struct {
	inner C.HederaSignature
}

// Generate a new [SecretKey] from a cryptographically secure pseudo-random number generator (CSPRNG).
func GenerateSecretKey() SecretKey {
	return SecretKey{C.hedera_secret_key_generate()}
}

// Parse a [HederaSecretKey] from a hex-encoded string.
func SecretKeyFromString(s string) (SecretKey, error) {
	var key C.HederaSecretKey
	err := C.hedera_secret_key_from_str(C.CString(s), &key)
	if err != 0 {
		return SecretKey{}, hederaError(err)
	}

	return SecretKey{key}, nil
}

// Format this [SecretKey] as a hex-encoded string of the secret key encoded with a PKCS #8 wrapper (
// defined in RFC 5208).
func (key SecretKey) String() string {
	bytes := C.hedera_secret_key_to_str(&key.inner)
	defer C.free(unsafe.Pointer(bytes))

	return C.GoString(bytes)
}

// Sign a message with this [SecretKey]
func(key SecretKey) Sign(message []byte) (Signature, error) {
	ptr := C.uchar(message[0])

	var signature C.HederaSignature
	err := C.hedera_secret_key_sign(&key.inner, &ptr, C.size_t(len(message)), &signature)

	if err != 0 {
		return Signature{}, hederaError(err)
	}

	return Signature{signature}, nil
}

// Derive a [PublicKey] from this [SecretKey].
func (key SecretKey) Public() PublicKey {
	return PublicKey{C.hedera_public_key_from_secret_key(&key.inner)}
}

// An ed25519 public key.
type PublicKey struct {
	inner C.HederaPublicKey
}

// Format this [PublicKey] as a hex-encoded string of the secret key encoded with a PKIX wrapper (
// defined in RFC 3280).
func (key PublicKey) String() string {
	bytes := C.hedera_public_key_to_str(&key.inner)
	defer C.free(unsafe.Pointer(bytes))

	return C.GoString(bytes)
}

// verify a message and it's [Signature] are were signed by the [SecretKey] associated with
// this [PublicKey]
func(key PublicKey) verify(message []byte, signature Signature) bool {
	ptr := C.uchar(message[0])

	verified := C.hedera_public_key_verify(&key.inner, &signature.inner, &ptr, C.size_t(len(message)))

	if verified != 0{
		return true
	}

	return false
}

// Parse a [HederaPublicKey] from a hex-encoded string.
func PublicKeyFromString(s string) (PublicKey, error) {
	var key C.HederaPublicKey
	err := C.hedera_public_key_from_str(C.CString(s), &key)
	if err != 0 {
		return PublicKey{}, hederaError(err)
	}

	return PublicKey{key}, nil
}
