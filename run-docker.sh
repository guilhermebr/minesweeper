#!/bin/bash

BUILD_IMAGE='tsuru/alpine-go:latest'
CONTAINER_PROJECT_PATH='/go/src/github.com/guilhermebr/minesweeper'
BUILD_CMD="go build --ldflags '-linkmode external -extldflags \"-static\"' -o build/minesweeper ./cmd"

set -x

docker run --rm -v ${PWD}:${CONTAINER_PROJECT_PATH} -w ${CONTAINER_PROJECT_PATH} -e CC=/usr/bin/gcc -e GOPATH=/go ${BUILD_IMAGE} sh -c "${BUILD_CMD}"
docker build -t guilhermebr/minesweeper .
docker run --rm -p 3000:3000 guilhermebr/minesweeper
