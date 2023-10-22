#!/bin/bash
echo "go build"
go mod tidy
go build -o go-admin-api main.go
chmod +x ./go-admin-api
echo "build go-admin-api success"