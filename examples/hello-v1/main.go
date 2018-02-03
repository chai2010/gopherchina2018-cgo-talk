// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("你好, GopherChina 2018!\n"))
}
