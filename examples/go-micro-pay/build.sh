#!/bin/bash

echo "i wanna go build $1"


echo "copy $1 config file starting"

cp "./config/$1.env.json" "./config/env.json"

echo "copy $1 config file done"

go env

go build -o go-micro-pay *.go