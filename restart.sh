#!/bin/bash
echo "go build"
go mod tidy
go build -o go-admin-api main.go
chmod +x ./go-admin-api
echo "kill go-admin-api service"
killall go-admin-api # kill go-admin service
nohup ./go-admin-api server -c=config/settings.dev.yml >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
echo "run go-admin-api success"
ps -aux | grep go-admin-api
