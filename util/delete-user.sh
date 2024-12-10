#!/bin/bash

if [ "$#" -ne 1 ]; then
    exit 1
fi

ID=$1

curl -X DELETE localhost:3737/user/${ID}
