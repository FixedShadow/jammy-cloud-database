package global

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/conf"
	"github.com/sirupsen/logrus"
)

var (
	CONF *conf.ServerConfig
	LOG  *logrus.Logger
)
