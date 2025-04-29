package router

import (
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
}

func (r *AccountRouter) InitRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("account")
	baseApi := &v1.BaseApi{}
	{
		baseRouter.POST(":projectId/create", baseApi.CreateAccount)
	}
}
