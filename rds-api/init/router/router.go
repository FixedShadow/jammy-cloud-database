package router

import "github.com/gin-gonic/gin"

var (
	Router *gin.Engine
)

type CommonRouter interface {
	InitRouter(Router *gin.RouterGroup)
}

func RouteGroups() []CommonRouter {
	return []CommonRouter{
		&InstanceRouter{},
		&ParamRouter{},
		&BackupRouter{},
	}
}

func Routers() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	Router = gin.Default()

	PrivateGroup := Router.Group("/api/v1")
	for _, router := range RouteGroups() {
		router.InitRouter(PrivateGroup)
	}
	return Router
}
