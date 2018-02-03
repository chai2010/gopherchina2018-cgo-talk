# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default:
	go tool cgo main_go.go hello_go.go

run:
	@go build -o a.out
	@./a.out
	-@rm a.out

debug:
	nm ./_obj/_cgo_.o

clean:
	-rm -rf _obj
