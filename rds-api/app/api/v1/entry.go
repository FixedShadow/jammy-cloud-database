package v1

import pb "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"

type BaseApi struct{}

type FrontendServer struct {
	MysqlInstanceManagementService pb.MySQLInstanceManagementService
}

var ApiServer = new(FrontendServer)
