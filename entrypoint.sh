#!/bin/bash

if [ $# != 1 ]; then
  echo Specify an image to download
  exit 1
fi

IMAGE=$1
DEST="/download"

mkdir -p "$DEST" || exit 1
/download.sh "$DEST" "$IMAGE" > /dev/null 2>&1 || exit 1
cd "$DEST"
tar cf "/data/$IMAGE.tar" . || exit 1
