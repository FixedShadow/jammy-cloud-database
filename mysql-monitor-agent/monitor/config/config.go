package config

import (
	error2 "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/error"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"os"
)

var DefaultMetricDeltaDataTimeInSecond = utils.CronJobTimeSecond

func loadConfig(confName string, conf interface{}) (interface{}, error) {
	pwd := logs.GetCurrentDirectory()
	file, err := os.Open(pwd + "/" + confName)
	if err != nil {
		logs.GetLogger().Error("Open monitor configuration file error", zap.Error(err))
		return nil, error2.Errors.NoConfigFileFound
	}
	defer file.Close()
	decoder := jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		logs.GetLogger().Error("Parsing monitor configuration file error", zap.Error(err))
		return nil, error2.Errors.ConfigFileValidationError
	}
	logs.GetLogger().Info("Successfully loaded monitor configuration file")
	return conf, nil
}
