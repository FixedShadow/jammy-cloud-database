package v1

import (
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) CreateDBInstance(c *gin.Context) {
	apiServer := new(FrontendApiServer)
	_ = apiServer
}

func (b *BaseApi) DeleteDBInstance(c *gin.Context) {}

func (b *BaseApi) RestartDBInstance(c *gin.Context) {}

func (b *BaseApi) GetDBInstanceInfo(c *gin.Context) {}
