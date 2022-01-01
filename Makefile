.DEFAULT_GOAL := all

PACKAGE := github.com/byxorna/nycmesh-tool
timestamp := $(shell date +%s)

.PHONY: codegen
codegen:
	echo Generating UISP CLI
	swagger generate cli -f ./spec/uisp_swagger.json --cli-app-name uisp --skip-validation -t generated/go/uisp/

.PHONY: go_build
go_build: 
	go build -o bin/ $(PACKAGE)
	@echo Now run bin/nycmesh-tool!

.PHONY: go
go: go_build

all: go .git/hooks/pre-commit

.git/hooks/pre-commit:
	pre-commit install --allow-missing-config

.PHONY: clean
clean:
	rm bin/*

.PHONY: pre_commit
pre_commit:
	pre-commit run

.PHONY: test
test: go_build go_test

.PHONY: go_test
go_test:
	go test -v $(PACKAGE)/...
