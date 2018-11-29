#ifndef HEDERA_KEY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
#define HEDERA_KEY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E

#include <stdint.h>
#include "hedera-error.h"

#ifdef __cplusplus
extern "C" {
#endif

// TODO: Define a Keypair structure. Would be more efficient for signing.

/// A signature signed with a HederaSecretKey.
typedef struct { uint8_t bytes[64]; } HederaSignature;

/// Parse a [HederaSignature] from a hex-encoded string.
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_signature_from_str(const char* s, HederaSignature* out);

/// Format a [HederaSignature as a hex-encoded string of the signature.
///
///
/// Returns ownership of the string. Must be freed with [free].
extern char* hedera_secret_key_to_str(HederaSignature*);

/// An EdDSA secret key.
typedef struct { uint8_t bytes[32]; } HederaSecretKey;

/// Generate a new [HederaSecretKey] from a cryptographically secure pseudo-random number generator (CSPRNG).
extern HederaSecretKey hedera_secret_key_generate();

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

/// sign a message with a [HederaSecretKey]
///
/// Returns [HEDERA_ERROR_SUCCESS] (0) on success or any other value on error. Use [hedera_error_message] to retrieve
/// a message for the error.
extern HederaError hedera_secret_key_sign(HederaSecretKey*, const uint8_t* message, size_t message_len, HederaSignature* out);

/// An ed25519 public key.
typedef struct { uint8_t bytes[32]; } HederaPublicKey;

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

/// Verify a message using a [HederaSignature] and the corresponding [HederaPublicKey].
///
///
/// Returns 1 if verification passed, otherwise returns 0
extern int8_t hedera_public_key_verify(HederaPublicKey* p, HederaSignature* s, const uint8_t* message, size_t message_len);

#ifdef __cplusplus
}
#endif

#endif // HEDERA_KEY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
