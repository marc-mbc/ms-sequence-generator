#!/bin/sh

# Data Containers

docker stop ms-sequence-generator-redis-config-data
docker stop ms-sequence-generator-redis-persistent-data

docker rm ms-sequence-generator-redis-config-data
docker rm ms-sequence-generator-redis-persistent-data

# Redis server

docker stop ms-sequence-generator-redis-server
docker rm ms-sequence-generator-redis-server