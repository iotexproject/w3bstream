FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=1

WORKDIR /go/src
COPY ./ ./

RUN apk add --no-cache gcc musl-dev

RUN cd ./cmd/prover && go build -o prover

FROM alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/cmd/prover/prover /go/bin/prover

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/prover"]
