// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
#include <errno.h>

static void seterrno(int v) {
	errno = v;
}

static void noreturn() {}
*/
import "C"
import "fmt"

func main() {
	_, err := C.seterrno(9527)
	fmt.Println(err)

	// Output:
	// errno 9527

	x, _ := C.noreturn()
	fmt.Printf("%#v\n", x)

	// Output:
	// main._Ctype_void{}
}
