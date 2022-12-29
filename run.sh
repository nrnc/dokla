#!/bin/sh
make _go_clean_bin
make _build
docker image rm -f nchukka/dokla:latest
docker build . -t nchukka/dokla:latest --no-cache
docker network create local-dev
docker compose -p dokla up -d
