APP_NAME := spli
SRC_DIR := .
GO := go

# Commands
all: build

build:
	@echo "Building the binary..."
	$(GO) build -o $(APP_NAME) $(SRC_DIR)

docs:
	@echo "Building docs"
	$(GO) build -o $(APP_NAME) $(SRC_DIR)
	@mkdir ./docs
	@rm -rf ./docs/*.md
	./$(APP_NAME) docs --path ./docs

test:
	@echo "Running tests..."
	$(GO) test ./... -v

clean:
	@echo "Cleaning up..."
	@rm -rf $(APP_NAME)

release :
	@echo "Cleaning up..."
	@rm -rf dist
	@goreleaser release --snapshot --clean

fmt:
	@echo "Formatting the code..."
	$(GO) fmt ./...

vet:
	@echo "Vet the code..."
	$(GO) vet ./...

lint:
	@echo "Linting the code..."
	@golint ./...
