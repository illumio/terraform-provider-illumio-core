TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
TF_INSTALL_PATH?=$${HOME}/.terraform.d/plugins/illumio/illumio-core
PKG_DISPLAY_NAME?=Illumio Core
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=illumio-core

TEST_PARALLELISM?=4

default: build

tools:
	go mod vendor

build: fmtcheck
	go install
	mkdir -p $(TF_INSTALL_PATH)
	go build -o $(TF_INSTALL_PATH)/terraform-provider-$(PKG_NAME)

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck _testacc sweep

_testacc:
	-TF_ACC=1 go test $(TEST) -v $(TESTARGS) -parallel=$(TEST_PARALLELISM) -timeout 120m

sweep:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -sweep=all

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./illumio-core"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

docs: fmtcheck
ifeq (, $(shell which tfplugindocs))
	$(error "tfplugindocs must be in PATH - see https://github.com/hashicorp/terraform-plugin-docs")
endif
	tfplugindocs generate --ignore-deprecated true --provider-name $(PKG_NAME) --rendered-provider-name $(PKG_DISPLAY_NAME)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc _testacc sweep vet fmt fmtcheck errcheck vendor-status test-compile docs website website-test tools