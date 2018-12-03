#!/bin/bash

rm hello-world
GOARCH="amd64" \
GOHOSTARCH="amd64" \
GOHOSTOS="linux" \
GOOS="linux" \
VERSION="v2" \
APP_NAME="sample"
PORT="5005"
go build hello-world.go
docker build -t repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest . 
docker tag repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}
docker push repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}

# docker run --rm --name sample repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}

# kubectl -n mini-stage1 deployment.yaml