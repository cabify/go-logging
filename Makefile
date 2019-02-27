.PHONY: help test fmt vet simplify errcheck staticcheck lint shellcheck fix-fmt build

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

##
### Code validation
validate: ## Run all validation targets: test, fmt, vet, lint (but no acceptance testing)
validate: fmt vet test lint shellcheck

test: ## Run tests for all go packages
	@scripts/test.sh

fmt: ## Run goimports on all packages, printing files that don't match code-format if any
	@scripts/fmt.sh

vet: ## Run vet on all packages (more info running `go doc cmd/vet`)
	@scripts/vet.sh

deps: ## Prefetch deps to ensure required versions are downloaded
	@scripts/deps.sh

simplify: ## Run gosimple on all packages
	@scripts/simplify.sh

errcheck: ## Run errcheck to find ignored errors
	@scripts/errcheck.sh

staticcheck: ## Run staticcheck on the codebase
	@scripts/staticcheck.sh

lint: ## Run lint on the codebase, printing any style errors
	@scripts/lint.sh

shellcheck:	## Lint shell scripts for potential errors
	@scripts/shellcheck.sh

fix-fmt: ## Run goimports on all packages, fix files that don't match code-style
	@scripts/fix-fmt.sh

install-tools: ## Install the required tooling
	go get -u -v \
		golang.org/x/tools/cmd/goimports
