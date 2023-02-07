.PHONY: build
build:
	@go build ./...

.PHONY: clean-test-cache
clean-test-cache:
	@go clean -testcache

.PHONY: test
test: clean-test-cache
	@go test -v ./...

.PHONY: run
run:
	@go run ./...

.PHONY: lint
lint:
	@which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@golangci-lint run -v

.PHONY: goimports
goimports:
	@which goimports || go install golang.org/x/tools/cmd/goimports@latest
	@goimports -w .
