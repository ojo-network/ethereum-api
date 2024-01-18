# Build Image
FROM golang:1.20-alpine AS builder

ARG GH_ACCESS_TOKEN

RUN apk add --no-cache build-base git

WORKDIR /app

COPY go.mod go.sum ./
RUN git config --global url."https://${GH_ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
ENV export GOPRIVATE=github.com/ojo-network/indexer

RUN go mod download
COPY . .
RUN go build

## Final Image
FROM alpine

ENV HOME /app
WORKDIR $HOME
COPY --from=builder /app/ethereum-api /bin/ethereum-api
COPY --from=builder /app/sample-config.yaml /app/config.yaml

EXPOSE 5005

ENTRYPOINT [ "ethereum-api" ]
