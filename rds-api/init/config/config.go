package config

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/conf"
	"github.com/FixedShadow/jammy-cloud-database/rds-api/global"
	"gopkg.in/yaml.v3"
	"os"
)

func Init() {
	file, err := os.Open(global.ConfPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var config conf.ServerConfig
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	global.CONF = &config
}
