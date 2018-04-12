// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// GODEBUG=cgocheck=0 go run main.go

package main

/*
#include <stdint.h>

extern uintptr_t go_get_arr_ptr_v0(int* arr, int idx);
extern int*      go_get_arr_ptr_v1(int* arr, int idx);
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	values := []int32{42, 9, 101, 95}

	addr0 := C.go_get_arr_ptr_v0(
		(*C.int)(unsafe.Pointer(&values[0])),
		C.int(0),
	)
	fmt.Println("addr0:", addr0)

	// panic: runtime error: cgo result has Go pointer
	addr1 := C.go_get_arr_ptr_v1(
		(*C.int)(unsafe.Pointer(&values[0])),
		C.int(0),
	)
	fmt.Println("addr1:", uintptr(unsafe.Pointer(addr1)))
}

//export go_get_arr_ptr_v0
func go_get_arr_ptr_v0(arr *C.int, idx C.int) C.uintptr_t {
	base := C.uintptr_t(uintptr(unsafe.Pointer(arr)))
	p := base + C.uintptr_t(idx)*4
	return p
}

//export go_get_arr_ptr_v1
func go_get_arr_ptr_v1(arr *C.int, idx C.int) *C.int {
	base := uintptr(unsafe.Pointer(arr))
	p := (*C.int)(unsafe.Pointer(base + uintptr(idx)*4))
	return p
}
