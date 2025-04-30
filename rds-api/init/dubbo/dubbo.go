package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/app/api"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	mysqlinstancemanagement "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	sqlserverinstancemanagement "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/sqlserver"
)

func Init() {
	//init dubbo instance
	ins, err := dubbo.NewInstance(
		dubbo.WithName(global.CONF.DubboConfig.ClientName),
		dubbo.WithRegistry(
			registry.WithZookeeper(),
			registry.WithAddress(global.CONF.DubboConfig.RegistryAddress),
		),
	)
	if err != nil {
		panic(err)
	}
	cli, err := ins.NewClient()
	if err != nil {
		panic(err)
	}

	mysqlInstanceManagementService, err := mysqlinstancemanagement.NewMySQLInstanceManagementService(cli)
	sqlserverInstanceManagementService, err := sqlserverinstancemanagement.NewSQLServerInstanceManagementService(cli)
	if err != nil {
		panic(err)
	}
	api.ApiServer.MySQLInstanceManagementService = mysqlInstanceManagementService
	api.ApiServer.SQLServerInstanceManagementService = sqlserverInstanceManagementService
	//init other dubbo service here, like postgresqlInstanceManagementService
}
