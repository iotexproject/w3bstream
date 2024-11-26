FROM golang:1.22-alpine AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=1

WORKDIR /go/src
COPY ./ ./

RUN apk add --no-cache gcc musl-dev

RUN cd ./cmd/apinode && go build -o apinode

FROM alpine:3.20 AS runtime

ENV LANG en_US.UTF-8

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /go/src/cmd/apinode/apinode /go/bin/apinode
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/apinode"]
