.DEFAULT_GOAL := all

GOOS ?=
GOARCH ?=
PACKAGE := github.com/byxorna/nycmesh-tool
VERSION_PACKAGE := $(PACKAGE)/pkg/version
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_LDFLAGS := -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) -X "$(VERSION_PACKAGE).BuildDate=$(shell date -u)"
timestamp := $(shell date +%s)

.PHONY: codegen
codegen:
	echo Generating UISP CLI
	swagger generate cli -f ./spec/uisp_swagger.yaml --cli-app-name uisp --skip-validation -t generated/go/uisp/

.PHONY: go_build
go_build: 
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags='$(GO_LDFLAGS)' -o bin/nycmesh-tool $(PACKAGE)

.PHONY: go
go: go_build

all: go

.git/hooks/pre-commit:
	pre-commit install --allow-missing-config

.PHONY: clean
clean:
	rm bin/* || :

.PHONY: pre_commit
pre_commit:
	pre-commit run

.PHONY: test
test: go_build go_test

.PHONY: go_test
go_test:
	go test -v $(PACKAGE)/...
