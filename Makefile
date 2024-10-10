APP_NAME := spli
SRC_DIR := .
GO := go

# Commands
all: build

build:
	@echo "Building the binary..."
	$(GO) build -o $(APP_NAME) $(SRC_DIR)

test:
	@echo "Running tests..."
	$(GO) test ./... -v

clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME)

fmt:
	@echo "Formatting the code..."
	$(GO) fmt ./...

vet:
	@echo "Vet the code..."
	$(GO) vet ./...

lint:
	@echo "Linting the code..."
	@golint ./...
