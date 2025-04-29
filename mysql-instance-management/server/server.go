package server

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/global"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/handler"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/init/config"
	instancemanagement "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/proto"
	"github.com/dubbogo/gost/log/logger"
)

func Start() {
	config.Init()
	ins, err := dubbo.NewInstance(
		dubbo.WithName(global.CONF.Registry.ServiceName),
		dubbo.WithRegistry(
			registry.WithZookeeper(),
			registry.WithAddress(global.CONF.Registry.Address+":"+global.CONF.Registry.Port)),
		dubbo.WithProtocol(
			protocol.WithTriple(),
			protocol.WithPort(global.CONF.System.Port),
			protocol.WithIp(global.CONF.System.BindAddress),
		),
	)
	if err != nil {
		panic(err)
	}
	//server
	srv, err := ins.NewServer()
	if err != nil {
		panic(err)
	}
	if err := instancemanagement.RegisterMySQLInstanceManagementServiceHandler(srv, &handler.InstanceManagementService{}); err != nil {
		panic(err)
	}
	if err := srv.Serve(); err != nil {
		logger.Error(err)
	}
}
