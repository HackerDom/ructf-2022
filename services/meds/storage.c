#include <stdio.h>
#include <string.h>
#include <strings.h>

#include "storage.h"

struct tree_node {
	uuid_t key;
	char value[256];
};

struct tree_node tree[TREECAP];

bool is_empty(int node) {
	return tree[node].value[0] == 0;
}

void init_storage(const char* file_path) {
	bzero(tree, sizeof(tree));
}

int find_node(const uuid_t key, int root) {
	if (is_empty(root))
		return root;

	int result = uuid_compare(key, tree[root].key);

	if (result == 0)
		return root;

	int next = result < 0 ? 2 * root + 1 : 2 * root + 2;
	return find_node(key, next);
}

char * store_item(const uuid_t key, const char *value) {
	int node = find_node(key, 0);
	memcpy(tree[node].key, key, sizeof(uuid_t));
	return strcpy(tree[node].value, value);
}

char * load_item(const uuid_t key, char *buffer) {
	int node = find_node(key, 0);
	if (is_empty(node))
		return 0;
	return strcpy(buffer, tree[node].value);
}

void dump_tree(int id) {
	if (id >= TREECAP || is_empty(id))
		return;

    char uuid_str[37];
    uuid_unparse_lower(tree[id].key, uuid_str);
    uuid_str[7] = 0;

    printf("{text:{name:\"%s\"},children:[", uuid_str);
    dump_tree(2 * id + 1);
    dump_tree(2 * id + 2);
    printf("]},");
}

int get_tree_height(int id) {
	if (id >= TREECAP || is_empty(id))
		return 0;
	return 1 + max(get_tree_height(2 * id + 1), get_tree_height(2 * id + 2)); 
}