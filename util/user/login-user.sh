#!/bin/bash

if [ "$#" -ne 2 ]; then
    exit 1
fi

EMAIL=$1
PASS=$2

curl -X POST localhost:3737/user/login \
    -H "Content-Type: application/json" \
    -d "{\"email\": \"$EMAIL\", \"password\": \"$PASS\"}" -c - | grep token | awk '{print $7}'
