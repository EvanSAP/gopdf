#!/bin/bash

set -euxo

export GO111MODULE=on
export GOPATH=$HOME/go

go mod download

APP_NAME="gopdf"
PORT="5005"
VERSION="v9"

rm $APP_NAME | echo "File ${APP_NAME} does not exist. Don't care!"

CGO_ENABLED=0 GOARCH="amd64" GOHOSTARCH="amd64" GOHOSTOS="linux"  GOOS="linux" go build -installsuffix cgo $APP_NAME.go

GOPDF_BINARY_HASH=$(openssl dgst -md5 -binary gopdf  | xxd -pb)

docker build -t repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest .
docker tag repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${GOPDF_BINARY_HASH}
docker push repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${GOPDF_BINARY_HASH}

# docker run --rm --name sample repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}
# kubectl -n mini-stage1 apply -f deployment.yaml