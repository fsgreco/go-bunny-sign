.PHONY: build clean install test

build:
	go build -o bin/bunnysign ./cmd/bunnysign

clean:
	rm -rf bin/

install:
	go install ./cmd/bunnysign

test:
	go run ./cmd/bunnysign "Hello from make!"

.DEFAULT_GOAL := build