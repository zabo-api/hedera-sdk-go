#pragma once

 #include <stdint.h>
 #include <stdbool.h>
 #include "hedera-id.h"
 #include "hedera-timestamp.h"
 #include "hedera-array.h"

 #ifdef __cplusplus
 extern "C" {
 #endif

 typedef struct {
     HederaFileId file_id;
     int64_t size;

     HederaTimestamp expiration_time;
     bool deleted;

     HederaArray keys;

 } HederaFileInfo;

 #ifdef __cplusplus
 }
 #endif
