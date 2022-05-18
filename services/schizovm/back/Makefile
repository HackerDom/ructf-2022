SRC_DIR=./src
INC_DIR=./inc
EXECUTABLE=backend

CXX=clang++
LD=clang++

LDFLAGS=-lm -lpthread -lboost_thread -lboost_log -lboost_log_setup -lboost_system -lboost_program_options $(LINKER_FLAGS)
WARN_OPTS=-Wall -Werror -pedantic -Wno-deprecated-declarations

INCLUDES=-I$(INC_DIR)

COMPILER_FLAGS=

CXXFLAGS=$(WARN_OPTS) $(INCLUDES) $(COMPILER_FLAGS) \
		-std=c++2a -g -O0

LINK_EXECUTABLE=$(LD) $(LDFLAGS) -o $@ $^

COMPILE_CXX_SRC=$(CXX) $(CXXFLAGS) -c -o $@ $<

all: $(EXECUTABLE)

clean: clean_objects
	rm -rf $(EXECUTABLE)

clean_objects:
	rm -rf *.o

$(EXECUTABLE): main.o
	$(LINK_EXECUTABLE)

main.o: $(SRC_DIR)/main.cpp
	$(COMPILE_CXX_SRC)
