#pragma once

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    uint64_t seconds;
    uint32_t nanos;
} HederaDuration;

#ifdef __cplusplus
}
#endif
