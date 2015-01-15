# Author: https://github.com/jpoehls 
# Modified by: klim8d

.PHONY: build doc fmt lint run debug test vendor_clean vendor_get vendor_update vet

APP  = wab.io 
DEPS_FOLDER = .vendor
DEPS = github.com/garyburd/redigo/redis \
	   github.com/OneOfOne/xxhash/native

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GP := ${PWD}/$(DEPS_FOLDER):${GOPATH}
export GOPATH=$(GP)

default: build

build: vet
	go build -v -o ./bin/$(APP) ./src/

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

run:
	./bin/$(APP)

debug:
	go run ./src/main.go -debug=1

test:
	go test ./src/...

vendor_clean:
	rm -dRf ./$(DEPS_FOLDER)/src

# We have to set GOPATH to just the .vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/$(DEPS_FOLDER) go get -d -u -v \
    $(DEPS)

vendor_update: vendor_get
	rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .git` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .hg` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .bzr` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .svn`

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/...
