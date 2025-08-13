install:
	go get ./...

build:
	go build -o bin/ticker
	chmod +x bin/ticker

test:
	go test internal/