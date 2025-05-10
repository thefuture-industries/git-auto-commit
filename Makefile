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
	@echo "Running release build..."
	@go build -ldflags="-s -w" -trimpath -o bin/auto-commit .
	@upx.exe --best --lzma bin/auto-commit

test:
	@go test -v ./...

run: build
	@./bin/auto-commit
