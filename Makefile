# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY_NAME=main
TARGET=cmd

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) ./...

# Format code
fmt:
	$(GOFMT) ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the application
run: build
	./$(BINARY_NAME)

.PHONY: build test fmt clean run