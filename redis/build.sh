#!/bin/sh

# Data Containers

docker run --name ms-sequence-generator-redis-config-data \
    -v $(pwd)/config/redis.conf:/usr/local/etc/redis/redis.conf \
    debian:wheezy
docker run --name ms-sequence-generator-redis-persistent-data \
    -v $(pwd)/data:/data \
    debian:wheezy

# Redis server

docker run -d --name ms-sequence-generator-redis-server \
    --volumes-from ms-sequence-generator-redis-config-data \
    --volumes-from ms-sequence-generator-redis-persistent-data \
    -p 6379:6379 \
    redis:3.0.7


