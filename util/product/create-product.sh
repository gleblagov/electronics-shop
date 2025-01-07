#!/bin/bash

if [ "$#" -ne 5 ]; then
    exit 1
fi

TITLE=$1
PRICE=$2
QUANTITY=$3
CATEGORY=$4
RATING=$5

curl -X POST http://localhost:3737/product \
    -H "Content-Type: application/json" \
    -d "{\"title\": \"$TITLE\", \"price\": $PRICE, \"quantity\": $QUANTITY, \"category\": \"$CATEGORY\", \"rating\": $RATING}"
