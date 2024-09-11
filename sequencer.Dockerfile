# ghcr.io/machinefi/sequencer:latest
FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/sequencer && go build -o sequencer

FROM --platform=linux/amd64 alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/cmd/sequencer/sequencer /go/bin/sequencer
EXPOSE 9001

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/sequencer"]
