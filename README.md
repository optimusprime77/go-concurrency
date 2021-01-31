
## Description

Golang API to demonstrate optimization of concurrent HTTP requests.

## Requirements

* [Go](https://golang.org) >= 1.11

## Setup

1. `git clone git@github.com:optimusprime77/go-concurrency.git`
2. Set correct env vars in .env
3. `make install`

## Run

1. `make server`

Alternatively, you may run :

1. `docker build -t megatron/go-concurrency .`
2. `docker run -p 8000:8000 megatron/go-concurrency`


## API

### GET /api/v1/pigeon

## Test

1. `curl -v 127.0.0.1:8000/api/v1/pigeon`
2. `curl -v 127.0.0.1:8000/api/v1/pigeon?debug=true`
3. `curl -v 127.0.0.1:8000/api/v1/pigeon?debug=false`

Please feel free to use your favorite tool for testing concurrent requests to the same endpoint.