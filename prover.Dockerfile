# ghcr.io/machinefi/prover:latest
FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/prover && go build -o prover

FROM alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/cmd/prover/prover /go/bin/prover
EXPOSE 9002

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/prover"]
