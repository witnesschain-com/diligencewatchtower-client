#!/bin/sh

v=$(go version | { read _ _ v _; echo ${v#go}; })

if [ "$v" != "1.20" ]; then
    echo "go version not compatible, try experimental build.. exiting"
    exit 1
fi

go build -o watchtower -ldflags '-X main.VERSION=v1.0.1-internal'