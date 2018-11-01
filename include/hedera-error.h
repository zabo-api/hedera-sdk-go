#ifndef HEDERA_ERROR_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
#define HEDERA_ERROR_9999A0E8_2BD1_4C33_8071_D93A13B8A9E

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef int64_t HederaError;

// 0 = No error or successful
#define HEDERA_ERROR_SUCCESS 0

/// Return a message corresponding to the passing in error. Returns [NULL] if there is no corresponding error.
/// Error messages may only be obtained once.
extern char* hedera_error_message(HederaError);

#ifdef __cplusplus
}
#endif

#endif // HEDERA_ERROR_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
