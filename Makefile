.PHONY: fmt lint test

ifeq ($(OS),Windows_NT)
    RM := del /Q
    UPX := upx.exe
    BIN := bin/auto-commit
else
    RM := rm -f
    UPX := upx
    BIN := ./bin/auto-commit
endif

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run

check: fmt lint test
	@echo "All checks passed!"
	
build:
	@echo "Running build..."
	@go build -o $(BIN) ./cmd

buildrelease:
	@echo "Running release build (windows, linux)..."

	# windows
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit ./cmd
	upx.exe --best --lzma bin/auto-commit || true

	# linux
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-linux ./cmd
	upx --best --lzma bin/auto-commit-linux || true

buildrelease-update:
	@echo "Running release build update..."
	@go build -ldflags="-s -w" -trimpath -o $(BIN).update ./cmd
	$(UPX) --best --lzma $(BIN).update || true

test:
	@go test -v ./...

run: build
	@./bin/auto-commit
