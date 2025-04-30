package global

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/conf"
	"github.com/sirupsen/logrus"
)

var (
	CONF *conf.ServerConfig
	LOG  *logrus.Logger
)
