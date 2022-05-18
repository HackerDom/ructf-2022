#include "storage.h"

int main() {
	init_storage(0);

	char key[37];
	while (fgets(key, 37, stdin)) {
		uuid_t uuid;
		uuid_parse(key, uuid);

		store_item(uuid, "foo");
	}

	//printf("%d\n", get_tree_height(0));
	dump_tree(0);

	return 0;
}