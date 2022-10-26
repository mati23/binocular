#!/bin/bash

docker rm -f server-redis

docker run --name server-redis -p 6379:6379 -d redis

docker run -it --link server-redis:redis --rm redis redis-cli -h redis -p 6379
