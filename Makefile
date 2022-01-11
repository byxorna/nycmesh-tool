.DEFAULT_GOAL := all

GOOS ?=
GOARCH ?=
PACKAGE := github.com/byxorna/nycmesh-tool
VERSION_PACKAGE := $(PACKAGE)/pkg/version
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DESCRIBE := $(shell git describe --always HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GO_LDFLAGS := -X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) -X "$(VERSION_PACKAGE).BuildDate=$(shell date -u)" -X "$(VERSION_PACKAGE).Release=$(GIT_DESCRIBE)" -X "$(VERSION_PACKAGE).GitBranch=$(GIT_BRANCH)"
timestamp := $(shell date +%s)

.PHONY: codegen
codegen:
	echo Generating UISP CLI
	swagger generate cli -f ./spec/uisp_swagger.yaml --cli-app-name uisp --skip-validation --keep-spec-order -t generated/go/uisp/
	sed -i'' -e "s|github.com/byxorna/nycmesh-tool/models|github.com/byxorna/nycmesh-tool/generated/go/uisp/models|g" generated/go/uisp/**/*.go

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
	go test -v $(PACKAGE)/pkg/... $(PACKAGE)/generated/...
