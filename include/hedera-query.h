#ifndef HEDERA_QUERY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
#define HEDERA_QUERY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E

#include "hedera-error.h"
#include "hedera-client.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct HederaQuery HederaQuery;

extern HederaError hedera_query_cost(HederaQuery*, uint64_t*);

#ifdef __cplusplus
}
#endif

#endif // HEDERA_QUERY_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
