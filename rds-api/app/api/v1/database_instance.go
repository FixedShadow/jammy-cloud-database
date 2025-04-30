package v1

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1/helper"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/dto"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/service"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/constant"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	"github.com/gin-gonic/gin"
)

func (b *BaseApi) CreateDBInstance(c *gin.Context) {
	var req dto.DBInstanceSpec
	if err := c.ShouldBindJSON(&req); err != nil {
		global.LOG.Error(err)
		helper.ErrorWithMsg(c, constant.ParamError, constant.ErrTypeInvalidParams)
		return
	}
	instanceService := service.NewInstanceService()
	switch req.Engine {
	case "mysql":
		instanceService.CreateMySQLInstance(c, req)
	case "sqlserver":
		instanceService.CreateSQLServerInstance(c, req)
	case "postgresql":
		//
	default:
		helper.ErrorWithMsg(c, constant.ParamError, constant.ErrTypeInvalidParams)
	}
}

func (b *BaseApi) DeleteDBInstance(c *gin.Context) {}

func (b *BaseApi) RestartDBInstance(c *gin.Context) {}

func (b *BaseApi) GetDBInstanceInfo(c *gin.Context) {}
