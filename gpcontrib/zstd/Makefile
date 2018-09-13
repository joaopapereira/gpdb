MODULE = zstd/lib/libzstd
PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)

all:
	$(MAKE) -C zstd

install: all

