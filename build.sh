#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bi-proxy .

docker build -t dev.k2data.com.cn:5001/qqy/bi-proxy:dev-1.0 .