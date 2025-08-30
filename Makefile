.PHONY: build clean install #test

build:
    go build -o bin/bunnysign ./cmd/bunny-sign

clean:
    rm -rf bin/

install:
    go install ./cmd/bunny-sign

#test:
#    go test ./...

run:
    go run ./cmd/bunny-sign

.DEFAULT_GOAL := build