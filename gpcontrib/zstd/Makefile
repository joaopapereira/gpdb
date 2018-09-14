MODULE = zstd/lib/libzstd
PG_CONFIG = pg_config

ifdef USE_PGXS
PGXS := $(shell pg_config --pgxs)
include $(PGXS)
else
top_builddir = ../..
include $(top_builddir)/src/Makefile.global
include $(top_srcdir)/contrib/contrib-global.mk
endif


ifeq "$(with_zstd)" "yes"
all:
	$(MAKE) -C zstd

install: all
endif
