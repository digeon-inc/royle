ENV_GITLABCI_FILE := .env.gitlabci
ENV_GITLABCI = $(shell cat $(ENV_GITLABCI_FILE))


# App Server
.PHONY: run-dev
run-dev:
	docker compose -f docker-compose.dev.yml up --build

.PHONY: run-test
run-test:
	docker compose -f docker-compose.test.yml up --build

.PHONY: destroy
destroy:
	docker compose -f docker-compose.dev.yml -f docker-compose.test.yml down --volumes --remove-orphans
	rm -rf mysql/mysql-data


# Tools
.PHONY: tools
tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.1 # go 1.19に対応するバージョン

# Lint, Format
.PHONY: lint
lint: tools
	golangci-lint run ./... --timeout=5m

.PHONY: format
format: tools
	golangci-lint run ./... --fix

.PHONY: test
test:
	docker compose -f docker-compose.test.yml run --rm test-app go test -v ./...

.PHONY: test-gitlab
test-gitlab:
	go test -v -cover ./... -coverprofile=coverage.txt && go tool cover -func coverage.txt
	go get github.com/boumenot/gocover-cobertura
	go run github.com/boumenot/gocover-cobertura < coverage.txt > coverage.xml

.PHONY: test-coverage
test-coverage:
	docker compose -f docker-compose.test.yml run --rm test-api  sh -c 'go test -v -cover ./... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html && go tool cover -func cover.out'

.PHONY: check
check:
	echo "called"
