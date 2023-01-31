FROM golang:1.18-alpine AS builder

WORKDIR /greenera

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o greenera main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /greenera

COPY --from=builder /greenera/greenera ./
COPY ./docker/local/greenera.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./greenera.sh"]
