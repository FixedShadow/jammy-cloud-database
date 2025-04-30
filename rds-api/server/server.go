package server

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/init/config"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/init/dubbo"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/init/log"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/init/router"
	pb "github.com/FixedShadow/jammy-cloud-database/rds-api/proto/mysql"
	"net"
	"net/http"
)

type FrontendApiServer struct {
	MysqlInstanceManagementService pb.MySQLInstanceManagementService
	//postgresqlInstanceManagementService
	//sqlserverInstanceManagementService
}

func Start() {
	config.Init()
	log.Init()
	dubbo.Init()
	rootRouter := router.Routers()
	server := &http.Server{
		Addr:    global.CONF.System.BindAddress + ":" + global.CONF.System.Port,
		Handler: rootRouter,
	}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		panic(err)
	}
	//init server
	if err := server.Serve(ln); err != nil {
		panic(err)
	}
}
