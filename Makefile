SRCS := $(shell find . -type f -name '*.go' | grep -v vendor)


default: up

.PHONY: up
up: .env
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: build
build: gomod test
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags '-d -w -s'

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: gomod
gomod: go.mod
	go env -w GO111MODULE=on
	go mod tidy

.PHONY: test
test: fmt
	go test ./...