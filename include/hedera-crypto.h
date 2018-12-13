#pragma once

#include <stdint.h>
#include <stddef.h>

#include "hedera-error.h"

#ifdef __cplusplus
extern "C" {
#endif

/// An ed25519 public key.
typedef struct { uint8_t bytes[32]; } HederaPublicKey;

/// An EdDSA secret key.
typedef struct { uint8_t bytes[32]; } HederaSecretKey;

/// An EdDSA signature.
typedef struct { uint8_t bytes[64]; } HederaSignature;

/// Parse a [HederaSignature] from a hex-encoded string.
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_signature_from_str(const char* s, HederaSignature* out);

/// Format a [HederaSignature as a hex-encoded string of the signature.
extern char* hedera_signature_to_str(HederaSignature*);

/// Generate a new [HederaSecretKey] with a BIP-39 mnemonic using a cryptographically
/// secure random number generator.
///
/// The [password] is required with the mnemonic to reproduce the secret key.
extern HederaSecretKey hedera_secret_key_generate(const char* password, const char** mnemonic);

/// Parse a [HederaSecretKey] from a hex-encoded string.
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_secret_key_from_str(const char* s, HederaSecretKey* out);

/// Format a [HederaSecretKey] as a hex-encoded string of the secret key encoded with a PKCS #8 wrapper (
/// defined in RFC 5208).
///
/// Returns ownership of the string. Must be freed with [free].
extern char* hedera_secret_key_to_str(HederaSecretKey*);

/// Derive a [HederaPublicKey] from a [HederaSecretKey].
extern HederaPublicKey hedera_public_key_from_secret_key(HederaSecretKey*);

/// Format a [HederaPublicKey] as a hex-encoded string of the secret key encoded with a PKIX wrapper (
/// defined in RFC 3280).
///
/// Returns ownership of the string. Must be freed with [free].
extern char* hedera_public_key_to_str(HederaPublicKey*);

/// Parse a [HederaPublicKey] from a hex-encoded string.
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_public_key_from_str(const char* s, HederaPublicKey* out);

/// Sign a message with a [HederaSecretKey].
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_crypto_sign(
    HederaSecretKey*, const uint8_t* message, size_t message_len, HederaSignature* out);

/// Verify a message using a [HederaSignature] and the corresponding [HederaPublicKey].
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_crypto_verify(
    HederaPublicKey* p, HederaSignature* s, const uint8_t* message, size_t message_len, uint8_t* out);

#ifdef __cplusplus
}
#endif
