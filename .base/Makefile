# Author: klim8d

.PHONY: build doc fmt lint run debug test benchmark vendor_clean vendor_get vendor_update vet docker_deploy docker_build docker_run

APP  = wab.io
DEPS_FOLDER = vendor
DEPS = github.com/garyburd/redigo/redis \
	   github.com/ugorji/go/codec \
	   github.com/satori/go.uuid \
	   github.com/zenazn/goji

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
APPFOLDER = $(shell echo ${PWD}|sed -e 's/\(.*\)\/.*/\1/'|tr -d " \t\n\r")
#APPFOLDER = $(shell echo ${PWD})
$(info VAR="$(APPFOLDER)")
#GP := ${PWD}/$(DEPS_FOLDER):${GOPATH}
GP := $(APPFOLDER)/$(DEPS_FOLDER):${GOPATH}
export GOPATH=$(GP)
LOGPATH := "../.logs/http.log"
ROOTPATH := `readlink -f ../$(APP)`

CPUS = 8

default: build

build:
ifneq ("$(wildcard main.go)", "")
	go build -v -race -o bin/$(APP)
else
	go build -v -race
endif

doc:
	godoc -http=:6060 -index

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	find -name '*.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go fmt

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	find -name '*.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 golint

run:
	./bin/$(APP) -mode=1 #>> $(LOGPATH)

debug:
ifneq ("$(wildcard bin/$(APP))", "")
	dlv debug bin/$(APP)
else
	dlv test
endif

test:
ifdef run
		find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -p $(CPUS) -parallel $(CPUS) -test.run $(run)
else
		find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -p $(CPUS) -parallel $(CPUS)
endif
	#go test -v ./...

benchmark:
ifdef run
	find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -run=XXX -bench=. -test.run $(run)
else
	find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -run=XXX -bench=.
endif

profile:
ifdef run
	find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -parallel $(CPUS) -test.run $(run)
else
	find -name '*_test.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go test -v -parallel $(CPUS)
endif

vendor_clean:
	rm -dRf ./$(DEPS_FOLDER)/

# We have to set GOPATH to just the .vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/$(DEPS_FOLDER) go get -d -u -v $(DEPS) \
	&& mv ./$(DEPS_FOLDER)/src/* ./$(DEPS_FOLDER)/ \
	&& rm -r ./${DEPS_FOLDER}/src

vendor_update: vendor_get
	rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .git` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .hg` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .bzr` \
    && rm -rf `find ./$(DEPS_FOLDER)/src -type d -name .svn`

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
vet:
	find -name '*.go' -not -path "./.*" -not -path "./vendor*" -printf '%h\n' | sort -u | xargs -n1 -P1 go vet

docker_build:
	docker run --rm -v ${PWD}:/src centurylink/golang-builder
