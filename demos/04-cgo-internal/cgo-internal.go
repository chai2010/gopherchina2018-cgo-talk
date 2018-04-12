// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go tool cgo cgo-internal.go

package main

//int sum(int a, int b) { return a+b; }
import "C"

func main() {
	C.sum(1, 2)
}
