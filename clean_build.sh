#!/bin/sh

cd ./redis
sh clean_build.sh
cd ..

docker stop ms-sequence-generator

docker rm ms-sequence-generator