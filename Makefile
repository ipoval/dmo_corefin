GO = go

.PHONY: build
build:
	$(GO) build ./...

.PHONY: test
test:
	$(GO) test ./...

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: imports
imports:
	goimports -w .

.PHONY: fmt
fmt:
	gofmt -w .
