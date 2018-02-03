// Created by cgo - DO NOT EDIT

package main

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_char int8

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


func _Cfunc_CString(s string) *_Ctype_char {
	p := _cgo_cmalloc(uint64(len(s)+1))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], s)
	pp[len(s)] = 0
	return (*_Ctype_char)(p)
}

//go:linkname _cgo_runtime_gostring runtime.gostring
func _cgo_runtime_gostring(*_Ctype_char) string

func _Cfunc_GoString(p *_Ctype_char) string {
	return _cgo_runtime_gostring(p)
}
//go:cgo_import_static _cgo_973df858ea43_Cfunc_SayHello_in_c
//go:linkname __cgofn__cgo_973df858ea43_Cfunc_SayHello_in_c _cgo_973df858ea43_Cfunc_SayHello_in_c
var __cgofn__cgo_973df858ea43_Cfunc_SayHello_in_c byte
var _cgo_973df858ea43_Cfunc_SayHello_in_c = unsafe.Pointer(&__cgofn__cgo_973df858ea43_Cfunc_SayHello_in_c)

//go:cgo_unsafe_args
func _Cfunc_SayHello_in_c(p0 *_Ctype_char) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_973df858ea43_Cfunc_SayHello_in_c, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_973df858ea43_Cfunc_SayHello_in_go
//go:linkname __cgofn__cgo_973df858ea43_Cfunc_SayHello_in_go _cgo_973df858ea43_Cfunc_SayHello_in_go
var __cgofn__cgo_973df858ea43_Cfunc_SayHello_in_go byte
var _cgo_973df858ea43_Cfunc_SayHello_in_go = unsafe.Pointer(&__cgofn__cgo_973df858ea43_Cfunc_SayHello_in_go)

//go:cgo_unsafe_args
func _Cfunc_SayHello_in_go(p0 *_Ctype_char) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_973df858ea43_Cfunc_SayHello_in_go, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_export_dynamic SayHello_in_go
//go:linkname _cgoexp_973df858ea43_SayHello_in_go _cgoexp_973df858ea43_SayHello_in_go
//go:cgo_export_static _cgoexp_973df858ea43_SayHello_in_go
//go:nosplit
//go:norace
func _cgoexp_973df858ea43_SayHello_in_go(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_973df858ea43_SayHello_in_go
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_973df858ea43_SayHello_in_go(p0 *_Ctype_char) {
	SayHello_in_go(p0)
}

//go:cgo_import_static _cgo_973df858ea43_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_973df858ea43_Cfunc__Cmalloc _cgo_973df858ea43_Cfunc__Cmalloc
var __cgofn__cgo_973df858ea43_Cfunc__Cmalloc byte
var _cgo_973df858ea43_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_973df858ea43_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_973df858ea43_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
