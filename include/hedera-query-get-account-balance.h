#ifndef HEDERA_QUERY_GET_ACCOUNT_BALANCE_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
#define HEDERA_QUERY_GET_ACCOUNT_BALANCE_9999A0E8_2BD1_4C33_8071_D93A13B8A9E

#include "hedera-account-id.h"
#include "hedera-query.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__get_account_balance__new(HederaClient*, HederaAccountId account);

extern HederaError hedera_query__get_account_balance__answer(HederaQuery*, uint64_t*);

extern HederaError hedera_query__get_account_balance__cost(HederaQuery*, uint64_t*);

#ifdef __cplusplus
}
#endif

#endif // HEDERA_QUERY_GET_ACCOUNT_BALANCE_9999A0E8_2BD1_4C33_8071_D93A13B8A9E
