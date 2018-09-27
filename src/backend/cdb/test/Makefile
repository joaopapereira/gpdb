subdir=src/backend/cdb
top_builddir=../../../..
include $(top_builddir)/src/Makefile.global

#TARGETS=cdbbufferedread \
#	cdbsrlz \
#	cdbdistributedsnapshot

CGREEN_TARGETS=cdbutil \
	cdbbufferedread \
	cdbsrlz \
	cdbdistributedsnapshot \
	cdbappendonlyxlog

#TARGETS += cdbappendonlyxlog

include $(top_builddir)/src/backend/mock.mk

cdbdistributedsnapshot.green: $(MOCK_DIR)/backend/access/transam/distributedlog_mock.o

cdbappendonlyxlog.green: \
	$(MOCK_DIR)/backend/storage/file/fd_mock.o \
	$(MOCK_DIR)/backend/access/transam/xlogutils_mock.o

cdbutil.green: \
	$(MOCK_DIR)/backend/utils/error/elog_mock.o \
	../cdbutil.o