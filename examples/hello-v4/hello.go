// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "C"
import "fmt"

//export SayHello
func SayHello(s *C.char) {
	fmt.Print(C.GoString(s))
}
