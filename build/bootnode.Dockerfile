FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/bootnode &&  go build -o bootnode

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/cmd/bootnode/bootnode /go/bin/bootnode
EXPOSE 8000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/bootnode"]
