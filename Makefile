.PHONY: build test test-v fmt lint clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

clean:
	rm -f *.test *.out
