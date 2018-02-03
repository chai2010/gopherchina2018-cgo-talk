// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//#include "hello.h"
import "C"
import "fmt"

//export SayHello_in_go
func SayHello_in_go(s *C.char) {
	fmt.Printf("SayHello_in_go: %s\n", C.GoString(s))
}
