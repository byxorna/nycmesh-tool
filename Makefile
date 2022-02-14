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

UISP_SWAGGER_JSON_FILE := spec/uisp_swagger_original.json
UISP_SWAGGER_MODIFIED_FILE := spec/_uisp_swagger_patched.yaml
UISP_SWAGGER_YAML_FILE := spec/uisp_swagger.yaml

.PHONY: $(UISP_SWAGGER_MODIFIED_FILE)
$(UISP_SWAGGER_MODIFIED_FILE):
	# requires yq installed - http://mikefarah.github.io/yq/
	cat $(UISP_SWAGGER_JSON_FILE) | yq -y > $(UISP_SWAGGER_MODIFIED_FILE)
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-1-1644866420-prerelease-removal.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-2-1644867315-format-date-to-date-time.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-3-1644868976-outage.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-4-1644869339-sfp-plus.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-5-1644871790-yaml-y-bool-quoting.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-6-1644874350-devices-id-services-params-fix.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-7-1644876262-quoting-issues.patch
	patch -u $(UISP_SWAGGER_MODIFIED_FILE) -i spec/patch-8-1644878967-conflicting-model-names-outages.patch

.PHONY: $(UISP_SWAGGER_YAML_FILE)
$(UISP_SWAGGER_YAML_FILE): $(UISP_SWAGGER_MODIFIED_FILE)
	# requires yamllint - https://yamllint.readthedocs.io/en/stable/quickstart.html#installing-yamllint
	mv $(UISP_SWAGGER_MODIFIED_FILE) $(UISP_SWAGGER_YAML_FILE)
	yamllint -c .yamllint.yaml $(UISP_SWAGGER_YAML_FILE)

.PHONY: codegen
codegen: $(UISP_SWAGGER_YAML_FILE)
	# requires `go-swagger` - https://github.com/go-swagger/go-swagger
	echo Generating UISP CLI
	find generated/go/uisp -type f -name '*.go' -delete
	swagger generate cli -f $(UISP_SWAGGER_YAML_FILE) --cli-app-name uisp --skip-validation --keep-spec-order -t generated/go/uisp/
	sed -i'' -e "s|github.com/byxorna/nycmesh-tool/models|github.com/byxorna/nycmesh-tool/generated/go/uisp/models|g" generated/go/uisp/**/*.go
	ln -sfn ../../../../internal/uisp/cli/custom.go generated/go/uisp/cli/custom.go

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
