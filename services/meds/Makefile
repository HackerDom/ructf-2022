meds: meds.* storage.* http.* diag.* *.html
	python3 embed.py > resources.h
	gcc-7 meds.c storage.c http.c diag.c -m32 -O2 -std=c11 -luuid -Wl,--strip-all -o bin/meds -DDEBUG_ON

test_storage: test_storage.* storage.*
	gcc-7 test_storage.c storage.c -m32 -O2 -std=c11 -luuid -Wl,--strip-all -o bin/test_storage -DDEBUG_ON