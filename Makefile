# Сборка утилиты
build:
	@go build -o unix_grep_lite cmd/main.go

# Запуск утилиты
grep:build
ifdef PATTERN
ifdef FLAGS
	@./unix_grep_lite $(FLAGS) $(PATTERN) $(INPUT_FILE)
else
	@./unix_grep_lite $(PATTERN) $(INPUT_FILE)
endif
else
	$(error Usage: make grep PATTERN="search" FLAGS="-n" INPUT_FILE="example.txt" (INPUT_FILE is option))
endif

# Тестирование
test:
	@go test -v ./internal/usecase

test-cover:
	@go test -v -covermode=atomic -coverprofile=coverage.out ./internal/usecase

# Качество кода
fmt:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	@golangci-lint run

# Справка
help:
	@echo "Unix Grep Lite - Available commands:"
	@echo ""
	@echo "Build command:"
	@echo "  build         - Build binary for current OS"
	@echo ""
	@echo "Run commands:"
	@echo "  grep          - Search pattern in file"
	@echo "                  Usage: make grep PATTERN=\"search\" FLAGS=\"-n\" INPUT_FILE=\"file.txt\""
	@echo "                  PATTERN is required, FLAGS and INPUT_FILE are optional"
	@echo "                  If INPUT_FILE empty, reads from stdin"
	@echo ""
	@echo "Test commands:"
	@echo "  test          - Run all tests"
	@echo "  test-cover    - Run tests with coverage"
	@echo ""
	@echo "Code quality:"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golangci-lint"

.PHONY: build grep test test-cover help fmt vet lint
