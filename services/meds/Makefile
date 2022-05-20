ARCH=-m64
FLAGS=$(ARCH) -O2 -std=c11 -luuid -Wl,--strip-all

meds: meds.* storage.* http.* diag.* *.html
	python3 embed.py > resources.h
	gcc-7 meds.c storage.c http.c diag.c $(FLAGS) -o bin/meds
	gcc-7 meds.c storage.c http.c diag.c $(FLAGS) -o bin/meds_debug -DDEBUG_ON

test_storage: test_storage.* storage.*
	gcc-7 test_storage.c storage.c $(FLAGS) -o bin/test_storage -DDEBUG_ON
