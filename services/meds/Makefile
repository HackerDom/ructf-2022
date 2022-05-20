ARCH=-m64
FLAGS=$(ARCH) -O2 -std=c11 -luuid -Wl,--strip-all

meds: meds.* storage.* http.* diag.* *.html
	python3 embed.py > resources.h
	g++-7 -c $(ARCH) farmhash.cc -o farmhash.o
	g++-7 -c $(ARCH) farmhash_adapter.cc -o farmhash_adapter.o
	gcc-7 meds.c storage.c http.c diag.c farmhash.o farmhash_adapter.o $(FLAGS) -o bin/meds
	gcc-7 meds.c storage.c http.c diag.c farmhash.o farmhash_adapter.o $(FLAGS) -o bin/meds_debug -DDEBUG_ON
	rm *.o

test_storage: test_storage.* storage.*
	gcc-7 test_storage.c storage.c $(FLAGS) -o bin/test_storage -DDEBUG_ON

brute: brute.*
	g++-7 -c $(ARCH) farmhash.cc -o farmhash.o
	g++-7 -c $(ARCH) farmhash_adapter.cc -o farmhash_adapter.o
	gcc-7 brute.c farmhash.o farmhash_adapter.o diag.c $(FLAGS) -o bin/brute