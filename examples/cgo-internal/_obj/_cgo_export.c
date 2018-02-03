/* Created by cgo - DO NOT EDIT. */
#include <stdlib.h>
#include "_cgo_export.h"

extern void crosscall2(void (*fn)(void *, int, __SIZE_TYPE__), void *, int, __SIZE_TYPE__);
extern __SIZE_TYPE__ _cgo_wait_runtime_init_done();
extern void _cgo_release_context(__SIZE_TYPE__);

extern char* _cgo_topofstack(void);
#define CGO_NO_SANITIZE_THREAD
#define _cgo_tsan_acquire()
#define _cgo_tsan_release()

extern void _cgoexp_973df858ea43_SayHello_in_go(void *, int, __SIZE_TYPE__);

CGO_NO_SANITIZE_THREAD
void SayHello_in_go(char* p0)
{
	__SIZE_TYPE__ _cgo_ctxt = _cgo_wait_runtime_init_done();
	struct {
		char* p0;
	} __attribute__((__packed__)) a;
	a.p0 = p0;
	_cgo_tsan_release();
	crosscall2(_cgoexp_973df858ea43_SayHello_in_go, &a, 8, _cgo_ctxt);
	_cgo_tsan_acquire();
	_cgo_release_context(_cgo_ctxt);
}

CGO_NO_SANITIZE_THREAD
void _cgo_973df858ea43_Cfunc__Cmalloc(void *v) {
	struct {
		unsigned long long p0;
		void *r1;
	} __attribute__((__packed__)) *a = v;
	void *ret;
	_cgo_tsan_acquire();
	ret = malloc(a->p0);
	if (ret == 0 && a->p0 == 0) {
		ret = malloc(1);
	}
	a->r1 = ret;
	_cgo_tsan_release();
}
