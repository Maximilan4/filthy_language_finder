GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEP=dep
BINARY_NAME=build/connor

all: build
build:
	$(DEP) ensure
	$(GOBUILD) -o $(BINARY_NAME)
clean:
	$(GOCLEAN) && rm $(BINARY_NAME)
test:
	$(GOTEST)
.PHONY: dep build clean
