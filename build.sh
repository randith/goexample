#!/usr/bin/env bash

set -e

SERVICE_NAME=pwhash
BUILDER_IMAGE_TAG=$SERVICE_NAME:builder

# compile binary
docker build -f Dockerfile.builder . -t $BUILDER_IMAGE_TAG

# extract binary
CONTAINER_ID=`docker create $SERVICE_NAME:builder`
docker container cp $CONTAINER_ID:/go/src/github.com/randith/goexample/cmd/pwhash bin/$SERVICE_NAME

# cleanup
docker rm $CONTAINER_ID
docker rmi $BUILDER_IMAGE_TAG

# create service image
docker build -t $SERVICE_NAME .
