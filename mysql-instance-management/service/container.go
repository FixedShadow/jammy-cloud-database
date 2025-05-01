package service

import (
	"context"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/auth"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/const"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/global"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/utils/container"
	"github.com/canonical/lxd/shared/api"
	"strconv"
)

type ContainerService struct{}

type IContainerService interface {
	CreateContainer(ctx context.Context, containerSpecs model.ContainerCreateSpecs) (*model.ContainerInfo, error)
}

func NewContainerService() IContainerService {
	return &ContainerService{}
}

func (i *ContainerService) CreateContainer(ctx context.Context, containerSpecs model.ContainerCreateSpecs) (*model.ContainerInfo, error) {
	client, err := container.NewClientWithAuth(global.CONF.ContainerZoneConfig.Address, auth.CertFile, auth.KeyFile)
	if err != nil {
		return nil, err
	}
	instance := api.InstancesPost{}
	instance.InstanceType = containerSpecs.ContainerType
	instance.Name = containerSpecs.ContainerName
	instance.Config = map[string]string{}
	instance.Config["limits.cpu"] = strconv.Itoa(containerSpecs.CpuNum)
	instance.Config["limits.memory"] = strconv.Itoa(containerSpecs.Memory) + constant.GB
	instance.Source = api.InstanceSource{
		Type:        "image",
		BaseImage:   containerSpecs.ImageId,
		Fingerprint: containerSpecs.ImageId,
	}
	instance.Devices = map[string]map[string]string{
		"root": {
			"path": "/",
			"pool": "default",
			"size": strconv.Itoa(containerSpecs.DiskSize) + constant.GB,
			"type": "disk",
		},
	}
	op, err := client.CreateInstance(instance)
	if err != nil {
		return nil, err
	}
	//wait for the operation to complete.
	err = op.WaitContext(ctx)
	//start the vm
	state := api.InstanceStatePut{
		Action:  constant.ACTION_START,
		Timeout: -1,
	}
	op, err = client.UpdateInstanceState(instance.Name, state, "")
	if err != nil {
		return nil, err
	}
	err = op.WaitContext(ctx)
	if err != nil {
		return nil, err
	}
	containerInfo := model.ContainerInfo{}
	containerInfo.ContainerName = instance.Name
	//TODO add another info.
	return &containerInfo, err
}
