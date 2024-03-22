# ghcr.io/machinefi/prover:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/prover && go build -o prover

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/prover/prover /go/bin/prover
COPY --from=builder /go/src/test/contract/Store.abi /go/bin/test/contract/Store.abi
EXPOSE 9002

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/prover"]
