package router

import (
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/gin-gonic/gin"
)

type LogRouter struct{}

func (r *LogRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("log")
	baseApi := &v1.BaseApi{}
	{
		baseRouter.GET(":projectId/get", baseApi.GetDatabaseLog)
	}
}
