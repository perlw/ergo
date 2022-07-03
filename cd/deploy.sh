#!/bin/bash
dockerid=$(docker ps -qf "ancestor=ergo")
gunzip -c images/ergo.tar.gz | docker load
if [ "$dockerid" != "" ]; then
  docker stop $dockerid
fi;
docker run --restart unless-stopped -d -p 8001:80 ergo:latest
docker container prune -f
docker image prune -af
