package service

import (
	"context"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/auth"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/global"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/utils/container"
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
	_ = client
	return nil, err
}
