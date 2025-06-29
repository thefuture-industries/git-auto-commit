.PHONY: fmt lint test

fmt:
	gofmt -w .
	goimports -w .

lint:
	golangci-lint run

check: fmt lint test
	@echo "All checks passed!"
build:
	@echo "Running build..."
	@go build -o bin/auto-commit .

buildrelease:
	@echo "Running release build (windows, linux)..."

	@go build -ldflags="-s -w" -trimpath -o bin/auto-commit .
	@upx.exe --best --lzma bin/auto-commit

	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o bin/auto-commit-linux .
	@upx.exe --best --lzma bin/auto-commit-linux

buildrelease-update:
	@echo "Running release build update..."
	@go build -ldflags="-s -w" -trimpath -o bin/auto-commit.update .
	@upx.exe --best --lzma bin/auto-commit.update

test:
	@go test -v ./...

run: build
	@./bin/auto-commit
