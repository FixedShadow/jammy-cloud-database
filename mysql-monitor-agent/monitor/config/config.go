package config

import (
	error2 "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/error"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"os"
)

var (
	monitorConfig                      *MonitorConfig
	DefaultMetricDeltaDataTimeInSecond = utils.CronJobTimeSecond
)

type Storage struct {
	KernelBucket    string `json:"kernel_bucket"`
	AgentBucket     string `json:"agent_bucket"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	KernelTarFile   string `json:"kernel_tar_file"`
	AgentTarFile    string `json:"agent_tar_file"`
	DownloadPath    string `json:"download_path"`
}

type MonitorConfig struct {
	Net      string  `json:"net"`
	Url      string  `json:"url"`
	Topic    string  `json:"topic"`
	Endpoint string  `json:"endpoint"`
	Storage  Storage `json:"storage"`
}

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

func GetMonitorConfig() *MonitorConfig {
	return monitorConfig
}

func init() {
	var ok bool
	monitorConf, err := loadConfig(utils.ConfMonitorName, monitorConfig)
	if err != nil {
		logs.GetLogger().Error("open config file error")
		panic(error2.Errors.OpenConfigFileError)
	}
	if monitorConfig, ok = monitorConf.(*MonitorConfig); !ok {
		logs.GetLogger().Error("type conversion error")
		panic(error2.Errors.CastTypeError)
	}
}
