# Build the app
FROM golang:1.15 as builder

WORKDIR /go/src/github.com/optimusprime77/go-concurrency
ADD . .

RUN apt-get update && \
    apt-get -y install gcc mono-mcs && \
    rm -rf /var/lib/apt/lists/*

RUN make build

RUN make test

FROM golang:1.15-alpine3.13

EXPOSE 8000

COPY --from=builder /go/src/github.com/optimusprime77/go-concurrency .

CMD ["./main"]