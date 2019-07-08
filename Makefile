
NAME=gone-alexa
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMODDOWNLOAD=$(GOCMD) mod download
BINARY_NAME=out/${NAME}
BINARY_UNIX=$(BINARY_NAME)_unix

all: clean deps test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	GO111MODULE=on $(GOMODDOWNLOAD)

build-linux:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

package:
	zip handler.zip -j $(BINARY_NAME)

package-linux:
	zip handler.zip -j $(BINARY_UNIX)
