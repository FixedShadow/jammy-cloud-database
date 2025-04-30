package handler

import (
	"context"
	"encoding/json"
	constant "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/const"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/model"
)
import pb "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/proto"

type InstanceManagementService struct{}

func (s *InstanceManagementService) CreateDBInstance(ctx context.Context, req *pb.CreateDBInstanceRequest) (res *pb.CreateDBInstanceResponse, err error) {
	containerCreateSpecs := model.ContainerCreateSpecs{}
	containerTemplate := map[string]map[string]int{}
	_ = json.Unmarshal(constant.ContainerTemplate, &containerTemplate)
	containerCreateSpecs.CpuNum = containerTemplate[req.InstanceClass]["cpu"]
	containerCreateSpecs.Memory = containerTemplate[req.InstanceClass]["memory"]
	return res, nil
}
