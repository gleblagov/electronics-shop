#!/bin/bash

if [ "$#" -ne 1 ]; then
    exit 1
fi

ID=$1

curl localhost:3737/cart/${ID}
