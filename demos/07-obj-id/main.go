// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.10

package main

//extern char* NewGoString(_GoString_ s);
//extern void PrintGoString(char* s);
//extern void FreeGoString(char* s);
import "C"

import (
	"fmt"
	"unsafe"
)

//export NewGoString
func NewGoString(s string) *C.char {
	id := NewObjectId(s)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export PrintGoString
func PrintGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	s := id.Get().(string)
	fmt.Println(s)
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

func main() {
	gs := "你好, QingCloud!"

	cs := C.NewGoString(gs)
	defer C.FreeGoString(cs)

	C.PrintGoString(cs)
}
