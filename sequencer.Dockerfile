# ghcr.io/machinefi/sequencer:latest
FROM --platform=linux/amd64 golang:1.22-alpine AS builder
#FROM --platform=linux/amd64 golang:1.22 AS builder

ENV GO111MODULE=on

# ca-certificates is required to call HTTPS endpoints.
# tzdata is required to time zone info.
RUN apk update && apk upgrade && apk add --no-cache ca-certificates tzdata musl-dev gcc && update-ca-certificates

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/sequencer && CGO_ENABLED=1 CGO_LDFLAGS='-L./lib/linux-x86_64 -lioConnectCore' go build -ldflags "-s -w -extldflags '-static'" -o sequencer

FROM --platform=linux/amd64 scratch AS runtime

ENV LANG en_US.UTF-8

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/cmd/sequencer/sequencer /go/bin/sequencer
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/sequencer"]
