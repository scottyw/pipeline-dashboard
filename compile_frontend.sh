#!/bin/bash

set -e

cd frontend; yarn build; cd ..
cp frontend/build/static/js/* public/js
cp frontend/build/static/css/* public/css

cp frontend/build/index.html public/index.html
