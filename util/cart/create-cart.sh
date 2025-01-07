#!/bin/bash

if [ "$#" -ne 1 ]; then
    exit 1
fi

USER_ID=$1

curl -X POST http://localhost:3737/cart \
    -H "Content-Type: application/json" \
    -d "{\"user_id\": $USER_ID}"
