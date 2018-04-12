// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"sort"
	"time"
	"unsafe"
)

func main() {
	// []float64 强制类型转换为 []int
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1, (1 << 20): 0}
	var b = []float64{4, 2, 5, 7, 2, 1, 88, 1, (1 << 20): 0}

	fmt.Println("len(a):", len(a))
	fmt.Println("len(b):", len(b))

	// 以int方式给float64排序 a
	bench("fast", func() {
		var p []int = ((*[1 << 28]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
		sort.Ints(p)
	})

	// 普通排序 b
	bench("slow", func() {
		sort.Float64s(b)
	})
}

func bench(name string, task func()) {
	start := time.Now()
	defer func() {
		fmt.Printf("%s: %v\n", name, time.Now().Sub(start))
	}()

	task()
	return
}
