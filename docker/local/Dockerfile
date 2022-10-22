FROM golang:1.18-alpine AS builder

WORKDIR /alveychain

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o alveychain main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /alveychain

COPY --from=builder /alveychain/alveychain ./
COPY ./docker/local/alveychain.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478

ENTRYPOINT ["./alveychain.sh"]
