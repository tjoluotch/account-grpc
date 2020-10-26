#!/bin/bash

IMG="rpc"

run(){
echo "building docker container..."
docker run \
--restart unless-stopped \
-p 2000:2000/tcp \
-d \
--env-file config.env \
--network microservice \
--name grpc-account \
rpc:0.0.1
}

build(){
echo "building docker image for grpc account service..."
docker build \
--rm \
-t ${IMG}:0.0.1 .
echo "docker image build for ${IMG} completed..."
}

if [[ $1 == "build" ]]
then
  build
elif [[ $1 == "run" ]]
then
    run
fi