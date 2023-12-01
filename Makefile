all: build

build:
	go build -o bin/aoc ./cmd/aoc

format:
	go fmt ./...

test: build
	ginkgo calendar
