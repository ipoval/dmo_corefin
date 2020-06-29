GO = go

.PHONY: build
build:
	$(GO) build ./...

.PHONY: test
test:
	$(GO) test -timeout 15s ./...

.PHONY: testv
testverbose:
	$(GO) test -v -cover -race -timeout 15s ./...

.PHONY: tidy
tidy:
	$(GO) mod tidy

.PHONY: imports
imports:
	goimports -w .

.PHONY: fmt
fmt:
	gofmt -w .
