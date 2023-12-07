# ghcr.io/machinefi/znode:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/znode && go build -o znode

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/znode/znode /go/bin/znode
COPY --from=builder /go/src/test/contract/Store.abi /go/bin/test/contract/Store.abi
EXPOSE 9002

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/znode"]
