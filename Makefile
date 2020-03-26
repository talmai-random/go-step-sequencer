# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=go-step-sequencer

mkfile_path := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
GOPATH=$(mkfile_path):$(mkfile_path)/submodules/

all: test build

build: deps
	GOPATH=$(GOPATH) $(GOBUILD) -o $(BINARY_NAME) -v

test:
	GOPATH=$(GOPATH) $(GOTEST) -v sequencer

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	git submodule update --init --recursive
