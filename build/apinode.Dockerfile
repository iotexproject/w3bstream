FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/apinode && go build -o apinode

FROM alpine:3.20 AS runtime

ENV LANG en_US.UTF-8

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /go/src/cmd/apinode/apinode /go/bin/apinode
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/apinode"]
