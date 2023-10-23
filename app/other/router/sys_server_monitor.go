package router

import (
	"go-admin/app/other/apis"
	"go-admin/common/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/nicelizhi/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysServerMonitorRouter)
}

// 需认证的路由代码
func registerSysServerMonitorRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ServerMonitor{}
	r := v1.Group("/server-monitor").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.ServerInfo)
	}
}
