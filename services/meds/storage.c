#include "storage.h"

struct tree_node {
	int my_node;
	uuid_t key;
	char value[256];
};

int16_t tree[TREE_MAXNODES];
struct tree_node tree_items[TREE_MAXITEMS];
int16_t next_item;

bool is_empty(int node) {
	return tree[node] == 0;
}

bool is_item_empty(struct tree_node* item) {
	return item->value[0] == 0;
}

int16_t allocate_item() {
	int16_t item = next_item;
	if (!is_item_empty(&tree_items[item])) {
		tree[tree_items[item].my_node] = 0;
		bzero(&tree_items[item], sizeof(struct tree_node));
	}
	next_item = (item + 1) % TREE_MAXITEMS;
	fflush(stdout);
	return item + 1;
}

struct tree_node* get_item(int index) {
	if (index < 1 || index > TREE_MAXITEMS)
		return 0;
	return &tree_items[index - 1];
}

void init_storage(const char* file_path) {
	bzero(tree, sizeof(tree));
	bzero(tree_items, sizeof(tree_items));
	next_item = 0;
}

int find_node(const uuid_t key, int root) {
	if (is_empty(root))
		return root;

	int result = uuid_compare(key, get_item(tree[root])->key);

	if (result == 0)
		return root;

	int next = result < 0 ? 2 * root + 1 : 2 * root + 2;
	return find_node(key, next);
}

char * store_item(const uuid_t key, const char *value) {
	int node = find_node(key, 0);

	if (is_empty(node))
		tree[node] = allocate_item();

	struct tree_node* item = get_item(tree[node]);
	item->my_node = node;
	memcpy(item->key, key, sizeof(uuid_t));
	return strcpy(item->value, value);
}

char * load_item(const uuid_t key, char *buffer) {
	int node = find_node(key, 0);
	if (is_empty(node))
		return 0;
	return strcpy(buffer, get_item(tree[node])->value);
}

void dump_tree(int id) {
	if (id >= TREE_MAXNODES || is_empty(id))
		return;

    char uuid_str[37];
    uuid_unparse_lower(get_item(tree[id])->key, uuid_str);
    uuid_str[7] = 0;

    printf("{text:{name:\"%s\"},children:[", uuid_str);
    dump_tree(2 * id + 1);
    dump_tree(2 * id + 2);
    printf("]},");
}

int get_tree_height(int id) {
	if (id >= TREE_MAXNODES || is_empty(id))
		return 0;
	return 1 + max(get_tree_height(2 * id + 1), get_tree_height(2 * id + 2)); 
}