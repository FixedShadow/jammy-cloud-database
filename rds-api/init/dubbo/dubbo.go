package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/registry"
	v1 "github.com/FixedShadow/jammy-cloud-database/rds-api/app/api/v1"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	mysqlinstancemanagement "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
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
	if err != nil {
		panic(err)
	}
	v1.ApiServer.MysqlInstanceManagementService = mysqlInstanceManagementService
	//init other dubbo service here, like postgresqlInstanceManagementService
}
