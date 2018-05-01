#!/bin/sh

dockerd &
for i in {0..30}
do
  docker info > /dev/null 2>&1
  if [ $? == 0 ]; then
    downloader $*
    exit 0
  fi
  sleep 1s
done

echo Unabled to start docker daemon. 1>&2
