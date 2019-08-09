#!/bin/bash

rm -rf build

mkdir build
mkdir build/static

go build -o ./build/server cmd/server/main.go

npm run build --prefix frontend

cp -r frontend/dist/* build/static/
