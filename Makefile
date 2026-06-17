.PHONY: build test test-v test-short fmt lint clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

test-short:
	go test -short ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

clean:
	rm -f *.test *.out
