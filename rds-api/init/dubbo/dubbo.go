package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	mysqlinstancemanagement "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/server"
)

func Init() {
	//init dubbo instance
	ins, err := dubbo.NewInstance(
		dubbo.WithName("rds-api"),
		dubbo.WithRegistry(
			registry.WithZookeeper(),
			registry.WithAddress("192.168.2.88:2181"),
		),
	)
	if err != nil {
		panic(err)
	}
	cli, err := ins.NewClient()
	if err != nil {
		panic(err)
	}
	global.ApiServer = new(server.FrontendApiServer)
	mysqlInstanceManagementService, err := mysqlinstancemanagement.NewMySQLInstanceManagementService(cli)
	if err != nil {
		panic(err)
	}
	global.ApiServer.MysqlInstanceManagementService = mysqlInstanceManagementService
	//init other dubbo service here, like postgresqlInstanceManagementService
}
