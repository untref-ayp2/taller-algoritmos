.PHONY: build test fmt clean

build:
	go build ./...

test:
	go test ./...

fmt:
	go fmt ./...

test-ejercicios:
	go test ./... 2>&1 | grep -E "^(ok|FAIL|---)" || true

clean:
	rm -f *.test *.out
