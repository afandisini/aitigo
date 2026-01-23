.PHONY: build install

build:
	go build -o aitigo ./cmd/aitigo

install:
	go install ./cmd/aitigo
