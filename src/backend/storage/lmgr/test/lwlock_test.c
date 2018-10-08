#include <stdarg.h>
#include <stddef.h>
#include <setjmp.h>
#include "cmockery.h"


void test__LWLockAcquire_when_no_more_memory_is_available_to_hold_locks_it_elog(void) {

}


int
main(int argc, char* argv[])
{
    cmockery_parse_arguments(argc, argv);

    const UnitTest tests[] = {
            unit_test(test__RemoveLocalLock_Null)
    };

    return run_tests(tests);
}
