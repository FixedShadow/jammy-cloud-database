package handler

import (
	"context"
	"encoding/json"
	constant "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/const"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/service"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/utils/common"
)
import pb "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/proto"

type InstanceManagementService struct{}

func (s *InstanceManagementService) CreateDBInstance(ctx context.Context, req *pb.CreateDBInstanceRequest) (res *pb.CreateDBInstanceResponse, err error) {
	containerCreateSpecs := model.ContainerCreateSpecs{}
	containerTemplate := map[string]map[string]int{}
	_ = json.Unmarshal(constant.ContainerTemplate, &containerTemplate)
	containerCreateSpecs.CpuNum = containerTemplate[req.InstanceClass]["cpu"]
	containerCreateSpecs.Memory = containerTemplate[req.InstanceClass]["memory"]
	containerCreateSpecs.DiskSize = int(req.InstanceStorageGB)
	containerCreateSpecs.ContainerType = constant.DEFAULT_CONTAINER_TYPE
	//log here.
	containerCreateSpecs.ContainerName = constant.IMAGE_TYPE_MYSQL + "_" + common.GenerateRandomStringLess32(10)
	containerCreateSpecs.ImageType = req.Engine + req.EngineVersion
	containerCreateSpecs.ImageId = service.NewImageService().GetImageIdByType(containerCreateSpecs.ImageType)
	containerInfo, err := service.NewContainerService().CreateContainer(ctx, containerCreateSpecs)
	_ = containerInfo
	return res, nil
}
