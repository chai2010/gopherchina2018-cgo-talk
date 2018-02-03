// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//#include "hello.h"
import "C"

func main() {
	C.SayHello_in_c(C.CString("Hello Golang!"))
	C.SayHello_in_go(C.CString("Hello Clang!"))
}
