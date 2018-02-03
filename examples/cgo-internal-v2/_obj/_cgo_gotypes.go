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
type _Ctype_double float64

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_import_static _cgo_6c7c0a6d4446_Cfunc_pow
//go:linkname __cgofn__cgo_6c7c0a6d4446_Cfunc_pow _cgo_6c7c0a6d4446_Cfunc_pow
var __cgofn__cgo_6c7c0a6d4446_Cfunc_pow byte
var _cgo_6c7c0a6d4446_Cfunc_pow = unsafe.Pointer(&__cgofn__cgo_6c7c0a6d4446_Cfunc_pow)

//go:cgo_unsafe_args
func _Cfunc_pow(p0 _Ctype_double, p1 _Ctype_double) (r1 _Ctype_double) {
	_cgo_runtime_cgocall(_cgo_6c7c0a6d4446_Cfunc_pow, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
