#!/usr/bin/env bash
echo "project : $1"
echo "version : $2"
if [[ -z "$1" ]] || [[ -z "$2" ]]
then
      echo "\$1[project] is NULL OR \$2[version] is NULL"
      exit 0
fi
export APPLICATION_VERSION=$2
docker stack deploy -c /root/server/docker-compose.$1.yml sillyhat