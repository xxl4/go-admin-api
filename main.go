package main

import (
	"go-admin/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 1.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @license.name MIT
// @license.url https://github.com/nicelizhi/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
