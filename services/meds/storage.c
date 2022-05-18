#include "storage.h"

struct tree_node {
	int my_node;
	uuid_t key;
	char value[256];
};

int16_t tree[TREE_MAXNODES];
struct tree_node tree_items[TREE_MAXITEMS];
int16_t next_item;
int items_count;

uuid_t recent_keys[TREE_MAXITEMS];
int most_recent_key;

char * _store_item(const uuid_t key, const char *value);

char *render_uuid(const uuid_t uuid) {
	char *s = malloc(37);
	uuid_unparse_lower(uuid, s);
	return s;
}

bool is_empty(int node) {
	return tree[node] == 0;
}

bool is_item_empty(struct tree_node* item) {
	return item->value[0] == 0;
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
	items_count = 0;
	most_recent_key = 0;
}

int find_node(const uuid_t key, int root) {
	if (root >= TREE_MAXNODES)
		DEBUG("!! trying to access OOB node %d\n", root);
	
	if (is_empty(root))
		return root;

	int result = uuid_compare(key, get_item(tree[root])->key);

	if (result == 0)
		return root;

	int next = result < 0 ? 2 * root + 1 : 2 * root + 2;
	return find_node(key, next);
}

void remove_children(int root, struct tree_node* children, int* children_count) {
	if (is_empty(root))
		return;

	struct tree_node* item = get_item(tree[root]);
	DEBUG("!! deleting item %d (node %d) ~ %s\n", tree[root], root, render_uuid(item->key));

	memcpy(&children[(*children_count)++], item, sizeof(struct tree_node));
	bzero(item, sizeof(struct tree_node));
	tree[root] = 0;
	items_count--;

	remove_children(2 * root + 1, children, children_count);
	remove_children(2 * root + 2, children, children_count);
}

void delete_item(int root) {
	struct tree_node children[TREE_MAXITEMS];
	int children_count = 0;

	remove_children(root, children, &children_count);

	for (int i = 1; i < children_count; i++)
		_store_item(children[i].key, children[i].value);
}

int16_t allocate_item() {
	int16_t item;
	for (int i = 0; i < TREE_MAXITEMS; i++) {
		item = (next_item + i) % TREE_MAXITEMS;
		if (is_item_empty(&tree_items[item]))
			break;
	}
	next_item = (item + 1) % TREE_MAXITEMS;
	items_count++;
	return item + 1;
}

int find_oldest_node() {
	for (int i = 0; i <= TREE_MAXITEMS; i++) {
		int idx = (most_recent_key + i) % TREE_MAXITEMS;
		int node = find_node(recent_keys[idx], 0);
		if (is_empty(node))
			continue;
		DEBUG("!! oldest node is %s, i = %d, mrk = %d\n", render_uuid(recent_keys[idx]), i, most_recent_key);
		return node;
	}
	perror("no node to delete");
	exit(1);
}

char * _store_item(const uuid_t key, const char *value) {
	int node = find_node(key, 0);

	if (is_empty(node)) {
		if (items_count == TREE_MAXITEMS) {
			delete_item(find_oldest_node());
			return _store_item(key, value);
		}
		tree[node] = allocate_item();
	}

	DEBUG("!! storing %s to item %d (node %d)\n", render_uuid(key), tree[node], node);

	struct tree_node* item = get_item(tree[node]);
	item->my_node = node;
	memcpy(item->key, key, sizeof(uuid_t));
	return strcpy(item->value, value);
}

char * store_item(const uuid_t key, const char *value) {
	_store_item(key, value);

	DEBUG("!! outer storing %s @ mrk = %d\n", render_uuid(key), most_recent_key);
	memcpy(recent_keys[most_recent_key], key, sizeof(uuid_t));
	most_recent_key = (most_recent_key + 1) % TREE_MAXITEMS;
}

char * load_item(const uuid_t key, char *buffer) {
	int node = find_node(key, 0);

	DEBUG("!! loading %s from item %d (node %d)\n", render_uuid(key), tree[node], node);

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