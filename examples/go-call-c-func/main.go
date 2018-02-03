// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
#include <errno.h>
#include <math.h>

static int add(int a, int b) {
	return a+b;
}
static void seterrno(int v) {
	errno = v;
}
static void noreturn() {}
*/
import "C"
import "fmt"

func main() {
	v0 := C.add(1, 1)
	fmt.Println(v0)

	v1, err := C.add(2, 2)
	fmt.Println(v1, err)

	_, err = C.seterrno(9527)
	fmt.Println(err)

	x, _ := C.noreturn()
	fmt.Printf("%#v\n", x)
}
