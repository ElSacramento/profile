.PHONY: all
all: build

bin:
	mkdir -p bin/

.PHONY: vendor
vendor:
	rm -rf vendor
	go mod vendor

.PHONY: build
build: bin
	go build -o bin/profile ./cmd/*.go

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
