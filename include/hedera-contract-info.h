#pragma once

 #include <stdint.h>
 #include "hedera-id.h"
 #include "hedera-crypto.h"
 #include "hedera-timestamp.h"
 #include "hedera-duration.h"

 #ifdef __cplusplus
 extern "C" {
 #endif

 typedef struct {
     HederaContractId contract;
     HederaAccountId account;

     char* contract_account_id;

     HederaPublicKey* admin_key;

     HederaTimestamp expiration_time;
     HederaDuration auto_renew_period;

     uint64_t storage;
 } HederaContractInfo;

 #ifdef __cplusplus
 }
 #endif
