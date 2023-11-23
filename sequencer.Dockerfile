# ghcr.io/machinefi/sequencer:latest
FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/sequencer && go build -o sequencer

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/sequencer/sequencer /go/bin/sequencer
EXPOSE 9000
EXPOSE 9001

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/sequencer"]
