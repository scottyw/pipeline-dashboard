#!/bin/bash

set -e

go get
go run main.go --no-cache

./build.sh $1 $2
