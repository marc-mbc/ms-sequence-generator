#!/bin/sh

# Prepare Redis service
cd ./redis
sh build.sh
cd ..

# Run microservice

docker run -d --name ms-sequence-generator \
    -p 8080:3000 \
    --link ms-sequence-generator-redis-server:db \
    marcmbc/ms-sequence-generator:release-0.0.0
