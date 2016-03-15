PACKAGES=$(shell go list ./... | grep -v /vendor/)

all: deps test lint

$(GOPATH)/bin/glide:
	go get github.com/Masterminds/glide

$(GOPATH)/bin/golint:
	go get github.com/golang/lint/golint

$(GOPATH)/bin/cover:
	go get golang.org/x/tools/cmd/cover

$(GOPATH)/bin/goveralls:
	go get github.com/mattn/goveralls

deps: $(GOPATH)/bin/glide \
	  $(GOPATH)/bin/golint \
	  $(GOPATH)/bin/cover \
	  $(GOPATH)/bin/goveralls
	glide install

build:
	go build -tags netgo --ldflags '-extldflags "-static" -w' ./cmd/...

lint:
	@for pkg in $(PACKAGES) ; do \
		golint -min_confidence 0.3 $$pkg; \
	done

test-and-cover:
	@for pkg in $(PACKAGES); do \
		go test -covermode=count -coverprofile=coverage_tmp.out $$pkg || exit 1; \
		tail -n +2 coverage_tmp.out >> coverage.out 2> /dev/null || exit 0; \
	done

.PHONY: build
