.DEFAULT_GOAL := all

# override this with https://10.70.76.21/nms/api-docs/swagger.json for example, if you are authenticated
PACKAGE := github.com/byxorna/nycmesh-tool
config_file := config/$(CONFIG).yaml
timestamp := $(shell date +%s)

.PHONY: codegen
codegen:
	echo hello

.PHONY: go_build
go_build: 
	go build -o bin/ $(PACKAGE)

.PHONY: go
go: go_build

all: go .git/hooks/pre-commit

.git/hooks/pre-commit:
	pre-commit install
.PHONY: clean
clean:
	rm .npm_fast_cache || :
	rm -rf node_modules || :

.PHONY: pre_commit
pre_commit:
	pre-commit run

dev: go
	UISP_AUTH_TOKEN=$$(cat authtoken) bin/nycmesh-tool experiment
