#!/bin/bash
set -o allexport; source .release.env; set +o allexport

env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o linux/arm64/${APPLICATION_NAME}
#env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o linux/amd64/${APPLICATION_NAME}
#env CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o darwin/arm64/${APPLICATION_NAME}
#env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o darwin/amd64/${APPLICATION_NAME}

sudo cp linux/arm64/${APPLICATION_NAME} /usr/local/bin
${APPLICATION_NAME} version