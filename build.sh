#!/usr/bin/env bash

echo 'build server'
GOOS=linux GOARCH=amd64 go build -o ./bin/ls-server-linux64 ./server/main.go

echo 'build local:mac'
GOOS=darwin GOARCH=amd64 go build -o ./bin/ls-local-darwin64 ./local/main.go