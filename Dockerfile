# iotexdev/w3bstream-mainnet-node:version
FROM golang:1.21 AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src
COPY ./ ./

RUN cd ./cmd/node && go build -o node

FROM golang:1.21 AS runtime

COPY --from=builder /go/src/cmd/node/node /go/bin/node
COPY --from=builder /go/src/test/data/risc0-project-config.json /go/bin/test/data/risc0-project-config.json
COPY --from=builder /go/src/test/contract/Store.abi /go/bin/test/contract/Store.abi
EXPOSE 9000

WORKDIR /go/bin
ENTRYPOINT ["/go/bin/node"]
