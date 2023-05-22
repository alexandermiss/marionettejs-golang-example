#!/bin/sh

export LINGUANSKI_IMAGE=$1:$2
export SERVER_URL="$3"

make deploy
