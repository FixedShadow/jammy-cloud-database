package global

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/conf"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/server"
)

var (
	CONF      *conf.ServerConfig
	ApiServer *server.FrontendApiServer
)
