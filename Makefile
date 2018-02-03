# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default:
	go run server.go


.PHONY: build
build:
	GOBIN=`pwd`/_bin go install ./examples/...

.PHONY: images
images:
	cd ./images && make

.PHONY: test
test:
	go fmt ./examples/...
	go vet ./examples/...
	go test ./examples/...

clean:
