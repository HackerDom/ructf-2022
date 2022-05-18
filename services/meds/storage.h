#pragma once

#include "types.h"

#define TREE_MAXITEMS 10
#define TREE_MAXNODES 200000

void init_storage(const char* file_path);

char * store_item(const uuid_t key, const char *value);
char * load_item(const uuid_t key, char *buffer);

void dump_tree(int id);
int get_tree_height(int id);