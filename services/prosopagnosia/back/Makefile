SRC_DIR=./src
INC_DIR=./inc
EXECUTABLE=prosopagnosia

CXX=clang++
LD=clang++

LDFLAGS=-lm -lpthread -lboost_thread -lboost_log -lboost_log_setup -lboost_system -lboost_program_options -lpq $(LINKER_FLAGS)
WARN_OPTS=-Wall -Werror -pedantic -Wno-deprecated-declarations

INCLUDES=-I$(INC_DIR) -I/usr/include/postgresql/

COMPILER_FLAGS=

CXXFLAGS=$(WARN_OPTS) $(INCLUDES) $(COMPILER_FLAGS) \
		-std=c++17 -g -O2

LINK_EXECUTABLE=$(LD) $(LDFLAGS) -o $@ $^

COMPILE_CXX_SRC=$(CXX) $(CXXFLAGS) -c -o $@ $<

all: $(EXECUTABLE)

clean: clean_objects
	rm -rf $(EXECUTABLE)

clean_objects:
	rm -rf *.o

$(EXECUTABLE): main.o fs.o pg_connection.o pg_connection_pool.o strings.o sha256.o defer.o options_parser.o demo_service.o
	$(LINK_EXECUTABLE)

main.o: $(SRC_DIR)/main.cpp
	$(COMPILE_CXX_SRC)

fs.o: $(SRC_DIR)/services/fs.cpp $(INC_DIR)/services/fs.h
	$(COMPILE_CXX_SRC)

pg_connection.o: $(SRC_DIR)/models/pg_connection.cpp $(INC_DIR)/models/pg_connection.h
	$(COMPILE_CXX_SRC)

pg_connection_pool.o: $(SRC_DIR)/models/pg_connection_pool.cpp $(INC_DIR)/models/pg_connection_pool.h
	$(COMPILE_CXX_SRC)

strings.o: $(SRC_DIR)/tools/strings.cpp $(INC_DIR)/tools/strings.h
	$(COMPILE_CXX_SRC)

sha256.o: $(SRC_DIR)/tools/sha256.cpp $(INC_DIR)/tools/sha256.h
	$(COMPILE_CXX_SRC)

defer.o: $(SRC_DIR)/tools/defer.cpp $(INC_DIR)/tools/defer.h
	$(COMPILE_CXX_SRC)

options_parser.o: $(SRC_DIR)/options/options_parser.cpp $(INC_DIR)/options/options_parser.h
	$(COMPILE_CXX_SRC)

demo_service.o: $(SRC_DIR)/services/demo_service.cpp $(INC_DIR)/services/demo_service.h
	$(COMPILE_CXX_SRC)
