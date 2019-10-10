#!/bin/bash

set -e
env GOOS=linux go build -o web cmd/web/main.go

cd frontend; yarn build; cd ..

mkdir -p public
cp -R frontend/build/* public/

timestamp = `date +%s`
docker build --no-cache . -f Dockerfile -t gcr.io/infracore/ci-dashboard:$(timestamp)
docker push gcr.io/infracore/ci-dashboard:$(timestamp)

rm web
