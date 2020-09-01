mkfile := $(abspath $(lastword $(MAKEFILE_LIST)))
dir := $(dir $(mkfile))

.PHONY: test
test:
	@go test --cover ./...

.PHONY: tidy
tidy:
	@go mod tidy
