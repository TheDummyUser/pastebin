
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTIDY=$(GOCMD) mod tidy
GORUN=$(GOCMD) run
BINARY_NAME=pastebin

all: build

build:
	mkdir -p bin
	$(GOBUILD) -o bin/$(BINARY_NAME) ./cmd/main.go

run:
	$(GOTIDY)
	$(GORUN) ./cmd/main.go

tidy:
	$(GOTIDY)

clean:
	$(GOCLEAN)
	rm -rf bin
