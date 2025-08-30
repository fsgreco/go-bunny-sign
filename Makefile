.PHONY: build clean install #test

build:
	go build -o bin/bunnysign ./cmd/bunnysign

clean:
	rm -rf bin/

install:
	go install ./cmd/bunnysign

#test:
#	go test ./...

run:
	go run ./cmd/bunnysign

.DEFAULT_GOAL := build