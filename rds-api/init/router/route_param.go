package router

import (
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/gin-gonic/gin"
)

type ParamRouter struct{}

func (r *ParamRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("param")
	baseApi := &v1.BaseApi{}
	{
		baseRouter.POST(":projectId/apply", baseApi.ApplyParam)
		baseRouter.PUT(":projectId/resume", baseApi.ResumeParam)
		baseRouter.GET(":projectId/info", baseApi.GetParam)
	}
}
