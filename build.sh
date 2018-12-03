#!/bin/bash

rm hello-world
GOARCH="amd64" \
GOHOSTARCH="amd64" \
GOHOSTOS="linux" \
GOOS="linux" \
go build hello-world.go

docker build -t repository.hybris.com:5005/upsurge/sample:latest .
docker push repository.hybris.com:5005/upsurge/sample:latest

#docker run --rm --name sample repository.hybris.com:5005/upsurge/sample:latest