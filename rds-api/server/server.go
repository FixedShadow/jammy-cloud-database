package server

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/init/router"
	"net"
	"net/http"
)

func Start() {
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
