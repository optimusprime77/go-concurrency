.PHONY: build install server test

## build: build the application
build:
	export GO111MODULE="on"; \
	go mod download; \
	go mod vendor; \
	CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o main cmd/server/main.go

## install: fetches go modules
install:
	export GO111MODULE="on"; \
	go mod tidy; \
	go mod download \

## server: runs the server with -race
server:
	export GO111MODULE="on"; \
	go run -race cmd/server/main.go

## test: runs tests
test:
	go test -race ./...

## help: prints help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
