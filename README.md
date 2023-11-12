# Ethereum API

Indexer for prices from ethereum trade pools

## Setup

Copy the sample-config.yaml to config.yaml and modify accordingly.

Run it locally with `go run main.go`

## Publish New Docker Image

`docker build . -t us-west4-docker.pkg.dev/ojo-network/ethereum-api`

`docker push us-west4-docker.pkg.dev/ojo-network/ethereum-api`
