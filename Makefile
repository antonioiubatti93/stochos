COVER := cover.out
TFLAGS = -race -coverprofile=bin/$(COVER)

PHONY: fmt
fmt:
	go mod tidy
	gofumpt -w .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	mkdir -p bin
	go test $(TFLAGS) ./...

.PHONY: cover
cover:
	go tool cover -html=$(COVER) -o $(COVER:.out=.html)
