GO    					:= go
PROJECT_NAME            ?= code-gen
TEST                    ?= $(shell go list ./... | grep -v '/vendor/')
TESTARGS                ?= -v -race

.PHONY: all
all: build

.PHONY: setup
setup:
	@echo ">> installing dependencies"
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0
	@$(GO) get -u "github.com/golang/tools/cmd/goimports"

.PHONY: test
test:
	@echo ">> running tests"
	@$(GO) test $(TEST) $(TESTARGS)

.PHONY: fmt
fmt:
	@find . -name "*.go" | xargs goimports -w
	@find . -name "*.go" | xargs gofmt -w

.PHONY: lint
lint:
	@echo ">> linting code"
	@golangci-lint run

.PHONY: build
build:
	@echo ">> building binaries"
	@$(GO) build -o code-gen

.PHONY: clean
clean:
	@echo ">> clean project"
	@rm -rf code-gen *.o
