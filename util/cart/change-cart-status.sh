#!/bin/bash

if [ "$#" -ne 2 ]; then
    exit 1
fi

ID=$1
STATUS=$2

curl -X PATCH localhost:3737/cart/${ID} \
    -H "Content-Type: application/json" \
    -d "{\"status\": \"$STATUS\"}"
