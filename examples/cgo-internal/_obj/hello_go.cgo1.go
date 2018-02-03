// Created by cgo - DO NOT EDIT

//line /Users/chai/go/src/github.com/chai2010/gopherchina2018-cgo-talk/examples/cgo-internal/hello_go.go:5
package main

//line /Users/chai/go/src/github.com/chai2010/gopherchina2018-cgo-talk/examples/cgo-internal/hello_go.go:9
import "fmt"

//line /Users/chai/go/src/github.com/chai2010/gopherchina2018-cgo-talk/examples/cgo-internal/hello_go.go:12
func SayHello_in_go(s *_Ctype_char) {
	fmt.Printf("SayHello_in_go: %s\n", _Cfunc_GoString(s))
}
