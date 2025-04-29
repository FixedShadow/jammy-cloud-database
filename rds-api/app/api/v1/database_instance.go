package v1

import (
	"context"
	"fmt"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/dto"
	mysqlinstancemanagement "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (b *BaseApi) CreateDBInstance(c *gin.Context) {
	var apiServer = ApiServer
	req := new(mysqlinstancemanagement.CreateDBInstanceRequest)
	instanceManagementRes, err := apiServer.MysqlInstanceManagementService.CreateDBInstance(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res := new(dto.CreateDBInstanceRes)
	res.Result = dto.CreateDBInstanceResult{
		InstanceId: instanceManagementRes.InstanceId,
	}
	c.JSON(http.StatusCreated, res)
}

func (b *BaseApi) DeleteDBInstance(c *gin.Context) {}

func (b *BaseApi) RestartDBInstance(c *gin.Context) {}

func (b *BaseApi) GetDBInstanceInfo(c *gin.Context) {}
