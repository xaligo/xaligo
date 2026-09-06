#ifndef XALIGO_ENGINE_H
#define XALIGO_ENGINE_H

#include <stddef.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define XALIGO_ENGINE_ABI_VERSION 5u

typedef struct XaligoEngineBuffer {
    uint8_t *data;
    size_t len;
    size_t capacity;
} XaligoEngineBuffer;

typedef struct XaligoEngineCancel XaligoEngineCancel;

uint32_t xaligo_engine_abi_version(void);

int32_t xaligo_engine_process(
    const uint8_t *input,
    size_t input_len,
    XaligoEngineBuffer *output
);

XaligoEngineCancel *xaligo_engine_cancel_new(void);
void xaligo_engine_cancel_set(XaligoEngineCancel *cancel);
void xaligo_engine_cancel_free(XaligoEngineCancel *cancel);
int32_t xaligo_engine_process_with_cancel(
    const uint8_t *input,
    size_t input_len,
    const XaligoEngineCancel *cancel,
    XaligoEngineBuffer *output
);

void xaligo_engine_buffer_free(XaligoEngineBuffer buffer);

#ifdef __cplusplus
}
#endif

#endif
