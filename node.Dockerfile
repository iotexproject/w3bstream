# ghcr.io/machinefi/node:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/node && go build -o node

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/node/node /go/bin/node
COPY --from=builder /go/src/test/contract/Store.abi /go/bin/test/contract/Store.abi
EXPOSE 9002

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/node"]
