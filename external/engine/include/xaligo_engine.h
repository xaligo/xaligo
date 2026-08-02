#ifndef XALIGO_ENGINE_H
#define XALIGO_ENGINE_H

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define XALIGO_ENGINE_ABI_VERSION 2u

typedef struct XaligoEngineBuffer {
    uint8_t *data;
    size_t len;
    size_t capacity;
} XaligoEngineBuffer;

uint32_t xaligo_engine_abi_version(void);

int32_t xaligo_engine_process(
    const uint8_t *input,
    size_t input_len,
    XaligoEngineBuffer *output
);

void xaligo_engine_buffer_free(XaligoEngineBuffer buffer);

#ifdef __cplusplus
}
#endif

#endif
