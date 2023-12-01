# ghcr.io/machinefi/coordinator:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/coordinator && go build -o coordinator

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/coordinator/coordinator /go/bin/coordinator
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/coordinator"]
