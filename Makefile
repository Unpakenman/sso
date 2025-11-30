.DEFAULT_GOAL := help

LOCAL_BIN=$(CURDIR)/bin


include bin-deps.mk

.PHONY: run
run: ## run project
	$ go run ./cmd/auth/main.go --use-local-env