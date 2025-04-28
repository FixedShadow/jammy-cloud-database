package router

import "github.com/gin-gonic/gin"

var (
	Router *gin.Engine
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router = gin.Default()
	return Router
}
