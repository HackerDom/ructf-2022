#include "storage.h"

#define QUEUE_SIZE (TREE_MAXITEMS * 10)
uuid_t queue[QUEUE_SIZE];
int queue_idx = 0;

int main(int argc, char** argv) {
	init_storage("data/storage");

	if (!strcmp(argv[1], "hack")) {

		char op[10];
		char key[37];
		char arg[30];
		while (scanf("%s %s %s\n", op, key, arg) != EOF) {
			if (strlen(key) != 36)
				continue;

			printf(".. %s %s %s\n", op, key, arg);

			uuid_t uuid;
			uuid_parse(key, uuid);

			char buf[256];
			if (!strcmp("store", op))
				store_item(uuid, arg);
			else if (!strcmp("load", op))
				printf("loaded: |%s|\n", load_item(uuid, buf));
			else
				printf("what? %s\n", op);
		}

		return 0;
	}

	{
		char key[37];
		while (fgets(key, 37, stdin)) {
			if (strlen(key) != 36)
				continue;

			uuid_t uuid;
			uuid_parse(key, uuid);

			DEBUG("!! store item: %s\n", key);

			if (queue_idx == QUEUE_SIZE) {
				printf("Too many items\n");
				return 1;
			}
			memcpy(queue[queue_idx++], uuid, sizeof(uuid_t));

			store_item(uuid, "foo");
		}
	}

	if (!strcmp(argv[1], "height")) {
		printf("%d\n", get_tree_height(0));
		return 0;
	}
	if (!strcmp(argv[1], "dump")) {
		dump_tree(0);
		return 0;
	}
	if (!strcmp(argv[1], "validate")) {
		int recent_items = min(TREE_MAXITEMS - PRELOAD_ITEMS, queue_idx);
		for (int i = 0; i < recent_items; i++) {
			DEBUG("!! load recent item #%d\n", i);
			char buf[256];
			char *val = load_item(queue[--queue_idx], buf);
			if (val == 0 || strcmp(val, "foo")) {
				printf("Item mismatch at index -%d\n", i);
				return 1;
			}
		}
		printf("All %d recent items loaded successfully.\n", recent_items);
		while (queue_idx > 0) {
			char buf[256];
			char *val = load_item(queue[--queue_idx], buf);
			if (val) {
				printf("A stale item was found at index %d\n", queue_idx);
				return 1;
			}
		}
		printf("No stale items were found.\n");

		return 0;
	}

	return 0;
}