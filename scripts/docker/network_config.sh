#!/bin/bash

OUTADDRSHOW=$(ip addr show | grep 206.8.3.1/24)

if [ ! -z "$OUTADDRSHOW" -a "$OUTADDRSHOW" != " " ]; then
    echo "network already created, initializing..."
else
    docker network create --subnet 206.8.3.0/24 password_gen_net_dev
fi
