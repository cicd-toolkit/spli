export GO15VENDOREXPERIMENT=1

exe = ./spli

.PHONY: all build install test coverage test-deps

all: install

test-deps:
	go get github.com/wadey/gocovmerge
	go get github.com/gucumber/gucumber/cmd/gucumber

build:
	go build

install:
	go install $(exe)

test-unit:
	go test ./cmd/ -v


test: test-unit

docker-build:
	sudo docker run --rm -v `pwd`:/go/src/github.com/cicd-toolkit/spli -w /go/src/github.com/cicd-toolkit/spli golang:1.6-alpine sh -c 'apk add --no-cache make git && make build'
	sudo docker build -t cicd-toolkit/spli .
	rm -f spli
