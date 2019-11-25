#!/usr/bin/env bash

if [[ "$1" == "image" ]]; then
    docker build -t locrep-maven .

    PORT=$port  BUILD_MODE=$mode \
    docker run -p $port:$port --name locrep-maven-container locrep-maven
else
    #run all tests
    BUILD_MODE=debug ginkgo -v -r

    go build -o locrep-maven

    PORT=$port BUILD_MODE=$mode  ./locrep-maven
fi