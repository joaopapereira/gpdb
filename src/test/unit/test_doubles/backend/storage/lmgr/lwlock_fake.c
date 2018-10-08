/*-------------------------------------------------------------------------
 *
 * lwlock.c
 *	  Lightweight lock manager Fake
 *
 *-------------------------------------------------------------------------
 */

#include "storage/spin.h"


void
LWLockAcquire(LWLockId lockid, LWLockMode mode) {

}

bool
LWLockConditionalAcquire(LWLockId lockid, LWLockMode mode) {

}


bool
LWLockAcquireOrWait(LWLockId lockid, LWLockMode mode) {

}


void
LWLockRelease(LWLockId lockid) {

}


void
LWLockReleaseAll(void) {

}