GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEP=dep
BINARY_NAME=build/connor

all: prepare build test
prepare:
	yes | cp -rf resources/ ~/connor_resources/

build:
	$(DEP) ensure
	$(GOBUILD) -o $(BINARY_NAME) -v
	$(GOBUILD) -o build/send_chats ./utils/test_send.go
	$(GOBUILD) -o build/read_results ./utils/test_read.go
clean:
	$(GOCLEAN) && rm $(BINARY_NAME)
test:
	$(GOTEST) -v
.PHONY: prepare build test clean
