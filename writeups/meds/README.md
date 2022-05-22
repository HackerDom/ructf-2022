# RuCTF 2022 | Meds

## Vulnerability

The service contains all elements in self-written [Binary Search Tree](https://en.wikipedia.org/wiki/Binary_search_tree).

- key: `GUID`, comes from client, controlled by attacker
- value: `tree_node` structure, contains data

```c
struct tree_node {
	int my_node;
	uuid_t key;
	bool protected;
	value_t value;
};
```

For optimization purposes, tree is stored as array:

- `left(i) = 2 * i`
- `right(i) = 2 * i + 1`

Where array element is an index inside another array of `tree_node` structures:

```c
struct {
	int16_t tree[TREE_MAXNODES];
	struct tree_node tree_items[TREE_MAXITEMS];
} data;
```

Again, the element in `tree` is an index of element in `tree_items`.

The tree is not balanced, so the attacker could create a long branch and perform a buffer out-of-bounds access.

## Exploitation

First of all, the tree is pre-filled. On initialization step all levels from root to 10-th are filled evenly:

```
        N/2
      /     \
   N/4       3N/4
  /  \       /   \
N/8  3N/8  5N/8  7N/8
```

Suppose we already have created the long branch, then our out-of-bounds node falls into the next `tree_items` array. We are interested in several functions of [storage.c](https://github.com/HackerDom/ructf-2022/blob/main/services/meds/storage.c):

1. `load_item()` is the part of storage interface

```c
char * load_item(const uuid_t key, value_t buffer) {
	int node = find_node(key, 0);

	DEBUG("!! loading %s from item %d (node %d)\n", render_uuid(key), data.tree[node], node);

	if (is_empty(node))
		return 0;
	return strcpy(buffer, get_item(data.tree[node])->value);
}
```

2. `find_node()` returns the index of element inside `data.tree` array, the element matches desired `key`

```c
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
```

3. `get_item()` return tree's value (`tree_node` structure from `tree_items` array)

```c
struct tree_node* get_item(int index) {
	if (index < 1 || index > TREE_MAXITEMS)
		return 0;
	return &data.tree_items[index - 1];
}
```

Again, more simplified version of `load_item()`:

```c
// tree_index is an index in `tree` array
int tree_index = find_node(key, 0);

// item_index is an index in `tree_items` array
int item_index = data.tree[idx];

// *node is an element of tree
tree_node *node = &data.tree_items[item_index - 1];

// value is a value of tree (corresponded to key)
value_t value = node->value;
```

We can select `target_index` in `data.tree` which we want to download and do the following:

1. Select an `controlled_index` of value which we can control
2. Make a long branch and do out-of-bounds write into `tree_items` array
3. Overwrite `data.tree[controlled_index]` with `target_index`, using oob-write
4. Read `data.tree[controlled_index]`, which will contain value of `target_index`

[Example sploit.py](/sploits/meds/sploit.py)
