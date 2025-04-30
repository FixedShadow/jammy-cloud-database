package service

import (
	"context"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/api"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1/helper"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/dto"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/constant"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	pbmysql "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/utils/common"
	"github.com/gin-gonic/gin"
)

type InstanceService struct{}

type InstanceServiceTraceId string

func NewInstanceService() *InstanceService {
	return &InstanceService{}
}

func (i InstanceService) CreateMySQLInstance(c *gin.Context, req dto.DBInstanceSpec) {
	createInstanceReq := &pbmysql.CreateDBInstanceRequest{}
	bindCreateInstanceParam(createInstanceReq, &req)
	traceId := common.Generate32RandomString()
	projectId := c.Param("projectId")
	//log traceId here.
	ctx := common.WithTraceId(context.Background(), traceId)
	ctx = context.WithValue(ctx, "projectId", projectId)
	instanceManagementService := api.ApiServer.MySQLInstanceManagementService
	createInstanceRes, err := instanceManagementService.CreateDBInstance(ctx, createInstanceReq)
	if err != nil {
		//log here.
		global.LOG.Error(err)
		helper.ErrorWithMsg(c, "", constant.ErrTypeInternalServer)
		return
	}
	instanceId := createInstanceRes.InstanceId
	retData := dto.CreateDBInstanceRes{}
	retData.RequestId = common.TraceId(ctx)
	retData.Result = dto.CreateDBInstanceResult{InstanceId: instanceId}
	helper.SuccessWithData(c, retData)
}

func (i InstanceService) CreateSQLServerInstance(c *gin.Context, req dto.DBInstanceSpec) {

}

func bindCreateInstanceParam(createInstanceReq *pbmysql.CreateDBInstanceRequest, req *dto.DBInstanceSpec) {
	createInstanceReq.InstanceName = req.InstanceName
	createInstanceReq.Engine = req.Engine
	createInstanceReq.EngineVersion = req.EngineVersion
	createInstanceReq.InstanceClass = req.InstanceClass
	createInstanceReq.InstanceStorageGB = req.InstanceStorageGB
	createInstanceReq.ParameterGroup = req.ParameterGroup
	createInstanceReq.InstanceStorageType = req.InstanceStorageType
	createInstanceReq.InstancePort = req.InstancePort
	createInstanceReq.StorageEncrypted = req.StorageEncrypted
	createInstanceReq.InstanceType = req.InstanceType
}
