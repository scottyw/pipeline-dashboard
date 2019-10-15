#!/bin/bash

set -e

echo "Running Go Build"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web cmd/web/main.go

cd frontend; yarn build; cd ..

mkdir -p public
cp -R frontend/build/* public/

TIMESTAMP=`date +%s`
docker build --no-cache . -f Dockerfile -t $1:$2


rm web
