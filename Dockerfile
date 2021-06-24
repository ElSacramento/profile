FROM golang:1.16 as builder

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make vendor build

FROM ubuntu:latest

COPY --from=builder /build/bin/* /usr/local/bin/
COPY --from=builder /build/storage/postgres/migrations /etc/migrations

ENTRYPOINT ["profile"]
