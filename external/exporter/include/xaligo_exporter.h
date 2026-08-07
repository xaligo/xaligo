#ifndef XALIGO_EXPORTER_H
#define XALIGO_EXPORTER_H

#include <stddef.h>
#include <stdint.h>

typedef struct XaligoExporterBuffer {
    uint8_t *data;
    size_t len;
    size_t capacity;
} XaligoExporterBuffer;

uint32_t xaligo_exporter_abi_version(void);
int32_t xaligo_exporter_process(const uint8_t *input, size_t input_len, XaligoExporterBuffer *output);
void xaligo_exporter_buffer_free(XaligoExporterBuffer buffer);

#endif
