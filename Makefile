# Lint, Format
.PHONY: lint
lint: tools
	golangci-lint run ./... --timeout=5m

.PHONY: format
format: tools
	golangci-lint run ./... --fix

# Test
.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -v -cover ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html && go tool cover -func cover.out

