#!/bin/bash
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
$env:GOOS='linux'
$env:GOARCH='amd64'
$env:CGO_ENABLED='0'
go build -o out/main main.go
& $env:GOPATH\bin\build-lambda-zip -o out/main.zip out/main
echo "build completed"