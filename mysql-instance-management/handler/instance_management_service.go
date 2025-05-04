package handler

import (
	"context"
	"encoding/json"
	constant "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/const"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/global"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/service"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/utils/common"
	"time"
)
import pb "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/proto"

type InstanceManagementService struct{}

func (s *InstanceManagementService) CreateDBInstance(ctx context.Context, req *pb.CreateDBInstanceRequest) (res *pb.CreateDBInstanceResponse, err error) {
	containerCreateSpecs := model.ContainerCreateSpecs{}
	containerCreateSpecs.ContainerName = constant.IMAGE_TYPE_MYSQL + "-" + common.GenerateRandomStringLess32(10)
	containerTemplate := map[string]map[string]int{}
	err = json.Unmarshal(constant.ContainerTemplate, &containerTemplate)
	if err != nil {
		return nil, err
	}
	containerCreateSpecs.CpuNum = containerTemplate[req.InstanceClass]["cpu"]
	containerCreateSpecs.Memory = containerTemplate[req.InstanceClass]["memory"]
	containerCreateSpecs.DiskSize = int(req.InstanceStorageGB)
	containerCreateSpecs.ContainerType = constant.DEFAULT_CONTAINER_TYPE
	//log here.
	containerCreateSpecs.ImageType = req.Engine + req.EngineVersion
	containerCreateSpecs.ImageId = service.NewImageService().GetImageIdByType(containerCreateSpecs.ImageType)
	/**
	It takes a lot of time to initialize the instance, so we use the async task.
	*/
	ctx2, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(global.CONF.ContainerZoneConfig.Timeout))
	go func(ctx context.Context) {
		defer cancel()
		containerInfo, err := service.NewContainerService().CreateContainer(ctx, containerCreateSpecs)
		if err != nil {
			global.LOG.Error("create container error: ", err.Error())
			return
		}
		err = service.NewContainerService().StartContainer(ctx, containerInfo)
		if err != nil {
			global.LOG.Error("start container error: ", err.Error())
			return
		}
	}(ctx2)
	res = new(pb.CreateDBInstanceResponse)
	res.InstanceId = containerCreateSpecs.ContainerName
	return res, nil
}
