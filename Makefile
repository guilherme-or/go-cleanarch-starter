# Scaffold commands
entity:
	@echo "Creating entity..."
	@go run ./cmd/scaffold/entity/main.go

usecase:
	@echo "Creating usecase..."
	@go run ./cmd/scaffold/usecase/main.go

# Testing coverage commands

go-test:
	@echo "Running all app tests..."
	@go test -v ./...

# Application commands

go-build:
	@echo "Building app..."
	@go build -o ./bin/app ./cmd/app/main.go

go-run:
	@echo "Running app..."
	@go run ./cmd/app/main.go

run: go-build
	@echo "Running app..."
	@./bin/app