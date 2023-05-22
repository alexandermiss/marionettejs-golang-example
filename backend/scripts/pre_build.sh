#!/bin/sh

mkdir ~/.docker
mv "$DOCKER_AUTH" ~/.docker/config.json
chmod g-r ~/.docker/config.json ; chmod o-r ~/.docker/config.json
