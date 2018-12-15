#pragma once

#include <stdint.h>
#include "hedera-id.h"
#include "hedera-query.h"
#include "hedera-file-info.h"

#ifdef __cplusplus
extern "C" {
#endif

extern HederaQuery* hedera_query__file_get_info__new(
    HederaClient*,
    HederaFileId file_id
);

extern HederaError hedera_query__file_get_info__get(
    HederaQuery*, 
    HederaFileInfo**
);

#ifdef __cplusplus
}
#endif