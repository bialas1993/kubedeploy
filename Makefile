TEST?=./...

default: test

build:
	go build -a -ldflags '-X "main.version=${version}" -X "main.date=${date}" -s -w' -o bin/kubedeploy
	
fmt: generate
	go fmt ./...

test: generate
	go get -t ./...
	go test $(TEST) $(TESTARGS)

generate:
	go generate ./...

.PHONY: default generate test
