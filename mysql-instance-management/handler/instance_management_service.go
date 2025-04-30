package handler

import (
	"context"
	constant "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/const"
)
import pb "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/proto"

type InstanceManagementService struct{}

func (s *InstanceManagementService) CreateDBInstance(ctx context.Context, req *pb.CreateDBInstanceRequest) (res *pb.CreateDBInstanceResponse, err error) {
	switch req.InstanceClass {
	case constant.MYSQL_SIZE_S1_MICRO:
	case constant.MYSQL_SIZE_S1_SMALL:
	}
	return res, nil
}
