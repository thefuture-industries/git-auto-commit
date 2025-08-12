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

	@find . -type f -name "*.sh" -exec sed -i 's/\r$$//' {} \;
	@go build -o $(BIN) ./cmd

buildrelease:
	@echo "Running release build (windows, linux)..."

	@find . -type f -name "*.sh" -exec sed -i 's/\r$$//' {} \;

	# windows
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-windows-amd64 ./cmd
	upx.exe --best --lzma bin/auto-commit-windows-amd64 || true

	# linux amd64
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-linux-amd64 ./cmd
	upx --best --lzma bin/auto-commit-linux-amd64 || true

	# linux arm64
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-linux-arm64 ./cmd
	upx --best --lzma bin/auto-commit-linux-arm64 || true

	# macOS (Intel)
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-macos-amd64 ./cmd

	# macOS (Apple Silicon / M1, M2)
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-macos-arm64 ./cmd

buildrelease-update:
	@echo "Running release build update..."
	
	@find . -type f -name "*.sh" -exec sed -i 's/\r$$//' {} \;

	@go build -ldflags="-s -w" -trimpath -o $(BIN).update ./cmd
	$(UPX) --best --lzma $(BIN).update || true

test:
	@go test -v ./...

run: build
	@./bin/auto-commit
