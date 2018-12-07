package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import "unsafe"

type SecretKey struct {
	inner C.HederaSecretKey
}

type PublicKey struct {
	inner C.HederaPublicKey
}

type Signature struct {
	inner C.HederaSignature
}

func GenerateSecretKey() SecretKey {
	return SecretKey{C.hedera_secret_key_generate()}
}

func SecretKeyFromString(s string) (SecretKey, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var key C.HederaSecretKey
	res := C.hedera_secret_key_from_str(cStr, &key)
	if res != 0 {
		return SecretKey{}, hederaLastError()
	}

	return SecretKey{key}, nil
}

func (secret SecretKey) Public() PublicKey {
	return PublicKey{C.hedera_public_key_from_secret_key(&secret.inner)}
}

func (secret SecretKey) String() string {
	return hederaString(C.hedera_secret_key_to_str(&secret.inner))
}

func (secret SecretKey) Sign(message []byte) (Signature, error) {
	ptr := C.CBytes(message)
	defer C.free(unsafe.Pointer(ptr))

	var signature C.HederaSignature
	res := C.hedera_crypto_sign(&secret.inner, (*C.uint8_t)(ptr), C.size_t(len(message)), &signature)
	if res != 0 {
		return Signature{}, hederaLastError()
	}

	return Signature{signature}, nil
}

func PublicKeyFromString(s string) (PublicKey, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var key C.HederaPublicKey
	res := C.hedera_public_key_from_str(cStr, &key)
	if res != 0 {
		return PublicKey{}, hederaLastError()
	}

	return PublicKey{key}, nil
}

func (public PublicKey) String() string {
	return hederaString(C.hedera_public_key_to_str(&public.inner))
}

func (public PublicKey) Verify(message []byte, signature Signature) (bool, error) {
	ptr := C.CBytes(message)
	defer C.free(unsafe.Pointer(ptr))

	var verified C.uint8_t
	res := C.hedera_crypto_verify(&public.inner, &signature.inner, (*C.uint8_t)(ptr), C.size_t(len(message)), &verified)
	if res != 0 {
		return false, hederaLastError()
	}

	return verified != 0, nil
}

func SignatureFromString(s string) (Signature, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var key C.HederaSignature
	res := C.hedera_signature_from_str(cStr, &key)
	if res != 0 {
		return Signature{}, hederaLastError()
	}

	return Signature{key}, nil
}

func (signature Signature) String() string {
	return hederaString(C.hedera_signature_to_str(&signature.inner))
}
