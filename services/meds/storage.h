#pragma once

#include "types.h"

#define TREE_MAXITEMS 11500
#define TREE_MAXNODES 10000000
#define PRELOAD_DEPTH 10
#define PRELOAD_ITEMS ((1 << (PRELOAD_DEPTH + 1)) - 1)

void init_storage(const char* file_path);

char * store_item(const uuid_t key, const value_t value);
char * load_item(const uuid_t key, value_t buffer);

void dump_tree(int id);
int get_tree_height(int id);

void generate_key(const value_t value, uuid_t key);
char *render_uuid(const uuid_t uuid); // debug only