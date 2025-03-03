# Scaffold commands

entity:
	@echo "Scaffolding entity..."
	@go run ./cmd/scaffold/entity/main.go

entity-test:
	@echo "Scaffolding entity test..."
	@go run ./cmd/scaffold/entity_test/main.go

repository:
	@echo "Scaffolding repository..."
	@go run ./cmd/scaffold/repository/main.go

dto:
	@echo "Scaffolding DTO..."
	@go run ./cmd/scaffold/dto/main.go

usecase:
	@echo "Scaffolding usecase..."
	@go run ./cmd/scaffold/usecase/main.go

# Testing coverage commands

test:
	@echo "Running all app tests..."
	@go test -v ./internal/...

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