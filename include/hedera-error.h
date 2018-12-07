#pragma once

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef int8_t HederaError;

/// 0 = No error.
#define HEDERA_ERROR_SUCCESS 0

/// 1 = There was an error. Use [hedera_last_error] to retrieve the error.
#define HEDERA_ERROR_FAILURE 1

/// Return the error message for the last error that occurred in this SDK.
/// Error messages may only be obtained once. Further attempts will return `NULL`.
/// Returns `NULL` if no error has occurred.
extern char* hedera_last_error();

#ifdef __cplusplus
}
#endif
