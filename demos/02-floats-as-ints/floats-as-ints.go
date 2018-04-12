// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	// []float64 强制类型转换为 []int
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)

	// 打印 []float64
	fmt.Println(a)
}
