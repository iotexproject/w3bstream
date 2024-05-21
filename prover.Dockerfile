# ghcr.io/machinefi/prover:latest
FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

RUN apk update && apk upgrade && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/prover && CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static'" -o prover

FROM scratch AS runtime

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/cmd/prover/prover /go/bin/prover
COPY --from=builder /go/src/test/contract/Store.abi /go/bin/test/contract/Store.abi
EXPOSE 9002

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/prover"]
