# ghcr.io/machinefi/sequencer:latest
FROM --platform=linux/amd64 golang:1.22 AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

#RUN cd ./cmd/sequencer && CGO_LDFLAGS='-L./lib/linux-x86_64 -lioConnectCore' go build -ldflags "-s -w -extldflags '-static'" -o sequencer
RUN cd ./cmd/sequencer && CGO_LDFLAGS='-L./lib/linux-x86_64 -lioConnectCore' go build -o sequencer

FROM --platform=linux/amd64 golang:1.22 AS runtime
#FROM --platform=linux/amd64 scratch AS runtime

COPY --from=builder /go/src/cmd/sequencer/sequencer /go/bin/sequencer
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/sequencer"]
