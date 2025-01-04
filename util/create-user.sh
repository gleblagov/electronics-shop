#!/bin/bash

if [ "$#" -ne 3 ]; then
    exit 1
fi

EMAIL=$1
PASS=$2
ROLE=$3

curl -X POST http://localhost:3737/user \
     -H "Content-Type: application/json" \
     -d "{\"email\": \"$EMAIL\", \"password\": \"$PASS\", \"role\": \"$ROLE\"}"

