.PHONY: build test run

all: build

deps:
	go get ./...

build: deps
	go build -o bin/main main.go

run:
	go run main.go
