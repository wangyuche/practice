#!/bin/bash
version=$2
filename=$1

docker rmi -f $(docker images | grep $filename)
docker build \
--build-arg version=$version \
--build-arg filename=$filename \
-t $filename:$version .

