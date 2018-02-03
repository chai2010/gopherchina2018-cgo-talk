// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
#include <stdio.h>
#include <stdint.h>

extern int* go_get_arr_ptr_v1(int *arr, int idx);
extern uintptr_t go_get_arr_ptr_v2(int *arr, int idx);

static void print_array_v1(int *arr, int n) {
	int i;
	for(i = 0; i < n; i++) {
		int *p = (int*)(go_get_arr_ptr_v1(arr, i));
		printf("%d ", *p);
	}
	printf("\n");
}
static void print_array_v2(int *arr, int n) {
	int i;
	for(i = 0; i < n; i++) {
		int *p = (int*)(go_get_arr_ptr_v2(arr, i));
		printf("%d ", *p);
	}
	printf("\n");
}
*/
import "C"
import "unsafe"

func main() {
	values := []int32{42, 9, 101, 95}

	C.print_array_v2(
		(*C.int)(unsafe.Pointer(&values[0])),
		C.int(len(values)),
	)
	C.print_array_v1(
		(*C.int)(unsafe.Pointer(&values[0])),
		C.int(len(values)),
	)
}

//export go_get_arr_ptr_v1
func go_get_arr_ptr_v1(arr *C.int, idx C.int) *C.int {
	base := uintptr(unsafe.Pointer(arr))
	p := (*C.int)(unsafe.Pointer(base + uintptr(idx)*4))
	return p
}

//export go_get_arr_ptr_v2
func go_get_arr_ptr_v2(arr *C.int, idx C.int) C.uintptr_t {
	base := C.uintptr_t(uintptr(unsafe.Pointer(arr)))
	p := base + C.uintptr_t(idx)*4
	return p
}
