#!/bin/sh

rsync -v scripts/deploy.sh $SERVER:$BASE_DIR/scripts
rsync -v Makefile $SERVER:$BASE_DIR
rsync -v docker-compose.yml $SERVER:$BASE_DIR
