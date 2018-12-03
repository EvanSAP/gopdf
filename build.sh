#!/bin/bash

export GO111MODULE=on
export GOPATH=$HOME/go

go get github.com/gin-gonic/gin
go get github.com/golang/protobuf

GOARCH="amd64" \
GOHOSTARCH="amd64" \
GOHOSTOS="linux" \
GOOS="linux" \
VERSION="v4" \
APP_NAME="gopdf"
PORT="5005"
rm $APP_NAME | echo "File ${APP_NAME} does not exist. Don't care!"
go build $APP_NAME.go
docker build -t repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest . 
docker tag repository.hybris.com:${PORT}/upsurge/${APP_NAME}:latest repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}
# docker push repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}

# docker run --rm --name sample repository.hybris.com:${PORT}/upsurge/${APP_NAME}:${VERSION}

# kubectl -n mini-stage1 deployment.yaml