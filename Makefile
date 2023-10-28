.PHONY: all build-mac build-linux clean

BINARY_NAME=chainpatrol

all: build-mac build-linux

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-mac main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux main.go

clean:
	rm -f bin/$(BINARY_NAME)-mac bin/$(BINARY_NAME)-linux

test:
	go test ./...
