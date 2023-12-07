# ghcr.io/machinefi/enode:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/enode && go build -o enode

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/enode/enode /go/bin/enode
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/enode"]
