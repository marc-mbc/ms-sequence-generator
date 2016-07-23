# ms-sequence-generator
API RestFul in Go to generate sequences

This microservice needs a Redis server.

You can Dockerize a Redis service:
```
# Data Containers
docker run --name ms-sequence-generator-redis-config-data \
    -v {path to your redis.conf in your host}:/usr/local/etc/redis/redis.conf \
    debian:wheezy
docker run --name ms-sequence-generator-redis-persistent-data \
    -v {path to a host folder to persist redis data}:/data \
    debian:wheezy

# Redis server
docker run -d --name ms-sequence-generator-redis-server \
    --volumes-from ms-sequence-generator-redis-config-data \
    --volumes-from ms-sequence-generator-redis-persistent-data \
    -p 6379:6379 \
    redis:3.0.7
```

And then run a ms-sequence-generator version:

```
docker run -d --name ms-sequence-generator \
    -p {set your host listening port}:3000 \
    --link ms-sequence-generator-redis-server:db \
    marcmbc/ms-sequence-generator:release-0.0.0
```

Or if you already have an existent redis server (listennig on port 6379) you can run a build without link a redis:

```
docker run -d --name ms-sequence-generator \
    -p {set your host listening port}:3000 \
    --add-host db:{your redis ip} \
    marcmbc/ms-sequence-generator:release-0.0.0
```

# Api Usage:

# Create a new sequence

The following request creates a new sequence called *Hotel* that will start at *1001*.

    curl -H "Accept: application/json" -H "Content-Type: application/json" -X POST -d '{"initialValue":1001}' http://localhost:8082/api/sequence/Hotel

# Get next number

The following three requests returns number *01001*, *01002* and *01003*.

    curl http://localhost:8082/sequence/Hotel/next
    { "number": "01001" }

    curl http://localhost:8082/sequence/Hotel/next
    { "number": "01002" }

    curl http://localhost:8082/sequence/Hotel/next
    { "number": "01003" }


# Get status of a sequence

The following request returns the status of the *Hotel* sequence without increasing the number.

    curl http://localhost:8082/sequence/Hotel
    { "number": "01003" }
