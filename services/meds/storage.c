#include "storage.h"
#include "farmhash_adapter.h"

enum store_flags {
	ST_NONE       = 0x00,
	ST_PROTECT    = 0x01, 
	ST_PERSIST    = 0x02, 
	ST_SAFE       = 0x04,
	ST_ADD_RECENT = 0x08
};

struct tree_node {
	int my_node;
	uuid_t key;
	bool protected;
	value_t value;
};

struct {
	int16_t tree[TREE_MAXNODES];
	struct tree_node tree_items[TREE_MAXITEMS];
} data;

int16_t next_item;
int items_count;

uuid_t recent_keys[TREE_MAXITEMS];
int most_recent_key;

const char* storage_path;

char *store_item_internal(const uuid_t key, const value_t value, enum store_flags flags);

char *render_uuid(const uuid_t uuid) {
	char *s = malloc(37);
	uuid_unparse_lower(uuid, s);
	return s;
}

bool is_empty(int node) {
	return data.tree[node] == 0;
}

bool is_item_empty(struct tree_node* item) {
	return item->value[0] == 0;
}

struct tree_node* get_item(int index) {
	if (index < 1 || index > TREE_MAXITEMS)
		return 0;
	return &data.tree_items[index - 1];
}

bool is_item_protected(struct tree_node* item) {
	return item->protected;
}

bool is_protected(int node) {
	return is_item_protected(get_item(data.tree[node]));
}

const uuid_t max_uuid = { 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff };

void get_average_uuid(const uuid_t a, const uuid_t b, uuid_t result) {
	memcpy(result, a, sizeof(uuid_t));
	uint32_t ha = reverse_bytes(*(uint32_t *)a);
	uint32_t hb = reverse_bytes(*(uint32_t *)b);
	uint32_t havg = ((uint64_t)ha + hb) / 2;
	*(uint32_t *)result = reverse_bytes(havg);
}

struct prefill_info {
	int depth;
	uuid_t lower;
	uuid_t upper;
};

void prefill_storage(int depth, const uuid_t lower, const uuid_t upper) {
	struct prefill_info queue[TREE_MAXITEMS];
	int front = 0;
	int back = 0;

	queue[back].depth = depth;
	bzero(queue[back].lower, sizeof(uuid_t));
	memcpy(queue[back].upper, max_uuid, sizeof(uuid_t));
	back++;

	while (front < back) {
		uuid_t avg;
		get_average_uuid(queue[front].lower, queue[front].upper, avg);
		store_item_internal(avg, "X", ST_PROTECT | ST_ADD_RECENT);

		if (queue[front].depth > 0) {
			if (back + 2 > TREE_MAXITEMS) {
				fprintf(stderr, "Tried to insert too many items\n");
				exit(1);
			}

			queue[back].depth = queue[front].depth - 1;
			memcpy(queue[back].lower, queue[front].lower, sizeof(uuid_t));
			memcpy(queue[back].upper, avg, sizeof(uuid_t));
			back++;

			queue[back].depth = queue[front].depth - 1;
			memcpy(queue[back].lower, avg, sizeof(uuid_t));
			memcpy(queue[back].upper, queue[front].upper, sizeof(uuid_t));
			back++;
		}

		front++;
	}
}


void save_to_file() {
	int fd;

	if ((fd = open(storage_path, O_WRONLY | O_CREAT, 0666)) < 0)
	{
		perror("Can't open storage\n");
		exit(1);
	}

	for (int i = 1; i <= TREE_MAXITEMS; i++) {
		struct tree_node* item = get_item(i);
		if (is_item_empty(item) || is_item_protected(item))
			continue;
		if (write(fd, item, sizeof(struct tree_node)) < 0)
		{
			perror("Failed to write item!");
			exit(1);
		}
	}

	close(fd);
}

void load_from_file() {
	int fd;

	if ((fd = open(storage_path, O_RDONLY, 0666)) < 0)
	{
		if (errno == ENOENT)
			return;
		perror("Can't open storage\n");
		exit(1);
	}

	int items_read = 0;
	struct tree_node item;
	while (true)
	{
		int bytesRead = read(fd, &item, sizeof(struct tree_node));

		if (!bytesRead)
			break;

		if (bytesRead < sizeof(item))
		{
			printf("Item %d is corrupt\n", items_read);
			break;
		}

		store_item_internal(item.key, item.value, ST_SAFE | ST_ADD_RECENT);
		items_read++;
	}

	printf("Loaded %d items.\n", items_count);

	close(fd);
}


void init_storage(const char* file_path) {
	bzero(data.tree, sizeof(data.tree));
	bzero(data.tree_items, sizeof(data.tree_items));
	next_item = 0;
	items_count = 0;
	most_recent_key = 0;
	storage_path = file_path;

	uuid_t lower;
	bzero(lower, sizeof(uuid_t));
	prefill_storage(10, lower, max_uuid);
	DEBUG("!! prefill inserted %d nodes!\n", items_count);

	load_from_file();
}

int find_node(const uuid_t key, int root) {
	if (root >= TREE_MAXNODES)
		DEBUG("!! trying to access OOB node %d\n", root);

	if (is_empty(root) || root >= TREE_MAXNODES)
		return root;

	int result = memcmp(key, get_item(data.tree[root])->key, sizeof(uuid_t));

	if (result == 0)
		return root;

	int next = result < 0 ? 2 * root + 1 : 2 * root + 2;
	return find_node(key, next);
}

void remove_children(int root, struct tree_node* children, int* children_count) {
	if (is_empty(root) || root >= TREE_MAXNODES)
		return;

	struct tree_node* item = get_item(data.tree[root]);
	DEBUG("!! deleting item %d (node %d) ~ %s\n", data.tree[root], root, render_uuid(item->key));

	memcpy(&children[(*children_count)++], item, sizeof(struct tree_node));
	bzero(item, sizeof(struct tree_node));
	data.tree[root] = 0;
	items_count--;

	remove_children(2 * root + 1, children, children_count);
	remove_children(2 * root + 2, children, children_count);
}

int child_comp(const void* a, const void *b) {
	return memcmp(((struct tree_node *)a)->key, ((struct tree_node *)b)->key, sizeof(uuid_t));
}

void delete_item(int root) {
	struct tree_node all_children[TREE_MAXITEMS];
	int children_count = 0;

	remove_children(root, all_children, &children_count);

	children_count--;
	struct tree_node* children = all_children + 1;

	if (children_count == 0)
		return;

	qsort(children, children_count, sizeof(struct tree_node), child_comp);

	int mid = children_count / 2;
	store_item_internal(children[mid].key, children[mid].value, ST_NONE);
	for (int i = mid + 1; i < children_count; i++) {
		store_item_internal(children[2 * mid - i].key, children[2 * mid - i].value, ST_NONE);
		store_item_internal(children[i].key, children[i].value, ST_NONE);
	}
	if (children_count % 2 == 0)
		store_item_internal(children[0].key, children[0].value, ST_NONE);
}

int16_t allocate_item() {
	int16_t item;
	for (int i = 0; i < TREE_MAXITEMS; i++) {
		item = (next_item + i) % TREE_MAXITEMS;
		if (is_item_empty(&data.tree_items[item]))
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
		// DEBUG("!! checking item = %d, node = %d, key = %s\n", idx, node, render_uuid(recent_keys[idx]));
		if (is_empty(node) || is_protected(node))
			continue;
		DEBUG("!! oldest node is %s, i = %d, mrk = %d\n", render_uuid(recent_keys[idx]), i, most_recent_key);
		return node;
	}
	fprintf(stderr, "No node to delete\n");
	exit(1);
}

char * store_item_internal(const uuid_t key, const value_t value, enum store_flags flags) {
	if (strlen(value) == 0)
		return 0;

	int node = find_node(key, 0);
	if (flags & ST_SAFE && node >= TREE_MAXNODES)
		return 0;

	if (is_empty(node)) {
		if (items_count == TREE_MAXITEMS) {
			delete_item(find_oldest_node());
			return store_item_internal(key, value, flags);
		}
		data.tree[node] = allocate_item();
	}

	DEBUG("!! storing %s to item %d (node %d), flags = %x\n", render_uuid(key), data.tree[node], node, flags);

	struct tree_node* item = get_item(data.tree[node]);
	item->my_node = node;
	item->protected |= (flags & ST_PROTECT) > 0;
	memcpy(item->key, key, sizeof(uuid_t));
	char* result = strcpy(item->value, value);

	if (flags & ST_ADD_RECENT) {
		DEBUG("!! outer storing %s @ mrk = %d\n", render_uuid(key), most_recent_key);
		memcpy(recent_keys[most_recent_key], key, sizeof(uuid_t));
		most_recent_key = (most_recent_key + 1) % TREE_MAXITEMS;
	}

	if (flags & ST_PERSIST)
		save_to_file();

	return result;
}

char * store_item(const uuid_t key, const value_t value) {
	return store_item_internal(key, value, ST_PERSIST | ST_ADD_RECENT);
}

char * load_item(const uuid_t key, value_t buffer) {
	int node = find_node(key, 0);

	DEBUG("!! loading %s from item %d (node %d)\n", render_uuid(key), data.tree[node], node);

	if (is_empty(node))
		return 0;
	return strcpy(buffer, get_item(data.tree[node])->value);
}

void dump_tree(int id) {
	if (id >= TREE_MAXNODES || is_empty(id))
		return;

    char uuid_str[37];
    uuid_unparse_lower(get_item(data.tree[id])->key, uuid_str);
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

void generate_key(const value_t value, uuid_t key) {
	fingerprint128(value, key);
}