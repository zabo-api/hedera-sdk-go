package hedera

import "C"

// #include <stdlib.h>
// #include "hedera-key.h"
import "C"
import (
	"unsafe"
)

// A Signature signed by a Secret Key
type Signature struct {
	inner C.HederaSignature
}

// Parse a [Signature] from a hex-encoded string.
func SignatureFromString(s string) (Signature, error) {
	var signature C.HederaSignature
	err := C.hedera_signature_from_str(C.CString(s), &signature)
	if err != 0 {
		return Signature{}, hederaError(err)
	}

	return Signature{signature}, nil
}

// Format this [Signature] as a hex-encoded string of the signature
func (signature Signature) String() string {
	bytes := C.hedera_signature_to_str(&signature.inner)
	defer C.free(unsafe.Pointer(bytes))

	return C.GoString(bytes)
}

// An EdDSA secret key.
type SecretKey struct {
	inner C.HederaSecretKey
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
func (key SecretKey) Sign(message []byte) (Signature, error) {
	ptr := C.CBytes(message)
	defer C.free(unsafe.Pointer(ptr))

	var signature C.HederaSignature
	err := C.hedera_secret_key_sign(&key.inner, (*C.uint8_t)(ptr), C.size_t(len(message)), &signature)

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

// verify a message and its [Signature] are were signed by the [SecretKey] associated with
// this [PublicKey]
func (key PublicKey) Verify(message []byte, signature Signature) (bool, error) {
	ptr := C.CBytes(message)
	defer C.free(unsafe.Pointer(ptr))

	var verified C.int8_t

	err := C.hedera_public_key_verify(&key.inner, &signature.inner, (*C.uint8_t)(ptr), C.size_t(len(message)), &verified)

	if err != 0 {
		return false, hederaError(err)
	}

	return verified != 0, nil
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
