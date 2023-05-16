EXEDIR := ./bin
COVER := $(EXEDIR)/cover.out
TFLAGS = -race -coverprofile=$(COVER)

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

.PHONY: build
build:
	mkdir -p $(EXEDIR)
	go build -o $(EXEDIR) ./cmd/main.go

run: build
	$(EXEDIR)/main -o $(EXEDIR)/paths.png
