package router

import (
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/gin-gonic/gin"
)

type BackupRouter struct{}

func (r *BackupRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("backup")
	baseApi := &v1.BaseApi{}
	{
		baseRouter.POST(":projectId/create", baseApi.CreateBackup)
		baseRouter.POST(":projectId/delete", baseApi.DeleteBackup)
	}
}
