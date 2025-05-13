# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
BINARY_NAME=hoyocheck
TARGET=cmd
DIR=bin


# Build the project
build:
	$(GOBUILD) -o ./$(DIR)/$(BINARY_NAME)

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
	./$(DIR)/$(BINARY_NAME)

.PHONY: build test fmt clean run