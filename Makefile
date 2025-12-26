.PHONY: build test run install clean

BINARY_NAME=togo

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test ./...

run:
	go run main.go

install:
	go install

clean:
	rm -f $(BINARY_NAME)
	go clean
