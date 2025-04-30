package api

import (
	pbmysql "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	pbsqlserver "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/sqlserver"
)


type FrontendServer struct {
	pbmysql.MySQLInstanceManagementService
	pbsqlserver.SQLServerInstanceManagementService
}

var ApiServer = new(FrontendServer)
