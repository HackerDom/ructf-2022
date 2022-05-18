#pragma once

#include "types.h"

void init_storage(const char* file_path);

char * store_item(const char *key, const char *value);
char * load_item(const char *key, char *buffer);
