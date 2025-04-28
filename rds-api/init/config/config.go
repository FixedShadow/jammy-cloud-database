package config

import (
	_ "embed"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/conf"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	"gopkg.in/yaml.v3"
)

//go:embed app.yaml
var AppYaml []byte

func Init() {
	var config conf.ServerConfig
	err := yaml.Unmarshal(AppYaml, &config)
	if err != nil {
		panic(err)
	}
	global.CONF = &config
}
