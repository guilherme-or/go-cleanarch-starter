# Scaffold commands

entity:
	@echo "Scaffolding entity..."
	@if [ -f ./bin/scaffold/entity ]; then ./bin/scaffold/entity; else go run ./cmd/scaffold/entity/main.go; fi

entity-test:
	@echo "Scaffolding entity test..."
	@if [ -f ./bin/scaffold/entity_test ]; then ./bin/scaffold/entity_test; else go run ./cmd/scaffold/entity_test/main.go; fi

repository:
	@echo "Scaffolding repository..."
	@if [ -f ./bin/scaffold/repository ]; then ./bin/scaffold/repository; else go run ./cmd/scaffold/repository/main.go; fi

dto:
	@echo "Scaffolding DTO..."
	@if [ -f ./bin/scaffold/dto ]; then ./bin/scaffold/dto; else go run ./cmd/scaffold/dto/main.go; fi

usecase:
	@echo "Scaffolding usecase..."
	@if [ -f ./bin/scaffold/usecase ]; then ./bin/scaffold/usecase; else go run ./cmd/scaffold/usecase/main.go; fi

# Testing coverage commands

test-entity:
	@echo "Running entity unit tests..."
	@go test -v ./internal/domain/entity/...

test:
	@echo "Running all app tests..."
	@go test -v ./internal/...

# Application commands

scaffold-build:
	@echo "Building scaffold app..."
	@go build -o ./bin/scaffold/entity ./cmd/scaffold/entity/main.go
	@go build -o ./bin/scaffold/entity_test ./cmd/scaffold/entity_test/main.go
	@go build -o ./bin/scaffold/repository ./cmd/scaffold/repository/main.go
	@go build -o ./bin/scaffold/dto ./cmd/scaffold/dto/main.go

go-build:
	@echo "Building app..."
	@go build -o ./bin/app ./cmd/app/main.go

go-run:
	@echo "Running app..."
	@go run ./cmd/app/main.go

build: scaffold-build go-build

run: go-build
	@echo "Running app..."
	@./bin/app