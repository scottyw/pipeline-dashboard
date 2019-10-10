#!/bin/bash

set -e
env GOOS=linux go build -o web cmd/web/main.go

cd frontend; yarn build; cd ..

mkdir -p public
cp -R frontend/build/* public/

docker build --no-cache . -f Dockerfile -t gcr.io/infracore/ci-dashboard:$(git rev-parse --short HEAD)
docker push gcr.io/infracore/ci-dashboard:$(git rev-parse --short HEAD)

rm web
