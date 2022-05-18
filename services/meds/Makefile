meds: meds.* storage.* http.*
	gcc-7 meds.c storage.c http.c -m32 -O2 -std=c11 -luuid -Wl,--strip-all -o bin/sesame

test_storage: test_storage.* storage.*
	gcc-7 test_storage.c storage.c -m32 -O2 -std=c11 -luuid -o bin/test_storage -DDEBUG_ON