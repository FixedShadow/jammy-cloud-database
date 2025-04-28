package router

import (
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/gin-gonic/gin"
)

type InstanceRouter struct{}

func (r *InstanceRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("instance")
	baseApi := &v1.BaseApi{}
	{
		baseRouter.POST("/:projectId/create", baseApi.CreateDatabase)
		baseRouter.DELETE("/:projectId/delete", baseApi.DeleteDatabase)
		baseRouter.POST("/:projectId/restart", baseApi.RestartDatabase)
		baseRouter.GET("/:projectId/info", baseApi.GetDatabaseInfo)
	}
}
