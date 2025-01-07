#!/bin/bash

if [ "$#" -ne 3 ]; then
    exit 1
fi

CART_ID=$1
PRODUCT_ID=$2
QUANTITY=$3

curl -X POST localhost:3737/cart/${CART_ID}/product \
    -H "Content-Type: application/json" \
    -d "{\"product_id\": $PRODUCT_ID, \"quantity\": $QUANTITY}"
