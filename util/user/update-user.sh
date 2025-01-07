#!/bin/bash

if [ "$#" -ne 4 ]; then
    exit 1
fi

ID=$1
EMAIL=$2
PASS=$3
ROLE=$4

curl -X PATCH http://localhost:3737/user/${ID} \
     -H "Content-Type: application/json" \
     -d "{\"email\": \"$EMAIL\", \"password\": \"$PASS\", \"role\": \"$ROLE\"}"

