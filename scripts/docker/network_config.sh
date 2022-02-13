#!/bin/bash

OUTADDRSHOW=$(ip addr show | grep 206.8.3.1/24)

if [[ OUTADDRSHOW != "" ]]; then
    docker network create --subnet 206.8.3.0/24 password_gen_net_dev
else
    echo "network already created, initializing..."
fi
