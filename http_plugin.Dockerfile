# ghcr.io/machinefi/http_plugin:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/http_plugin && go build -o http_plugin

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/http_plugin/http_plugin /go/bin/http_plugin
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/http_plugin"]
