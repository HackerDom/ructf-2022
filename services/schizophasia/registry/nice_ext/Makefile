MODULE_big = nice_ext
EXTENSION = nice_ext
OBJS = nice_ext.o
DATA = nice_ext--1.0.sql

SHLIB_LINK += -lz -lcrypto

PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)
