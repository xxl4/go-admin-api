package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 1.0.3
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @license.name Apache 2.0
// @license.url https://github.com/nicelizhi/go-admin-api/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
