package tpl

// MainTemplate main.go 内容
func MainTemplateGo() []byte {
	return []byte(`/*
	This is a demo name {{ .Name }} created by {{ .CreatorName }}
	{{ .CreatedTime }}
*/
package main

import	(
	"fmt"
)

func main() {
	fmt.Println("this is {{ .Name }}")
	// your code here
}
`)
}

// MakefileTemplate Makefile 内容
func MakefileTemplateGo() []byte {
	return []byte(`# This is a demo name {{ .Name }} created by {{ .CreatorName }}
# {{ .CreatedTime }}

GO    					:= go
PROJECT_NAME            ?= {{ .Name }}
TEST                    ?= $(shell go list ./... | grep -v '/vendor/')
TESTARGS                ?= -v -race

.PHONY: all
all: build run

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
	@$(GO) build -o $(PROJECT_NAME)

.PHONY: run
run:
	@echo ">> run binaries"
	@./$(PROJECT_NAME)

.PHONY: clean
clean:
	@echo ">> clean project"
	@rm -rf {{ .Name }} *.o
`)
}
