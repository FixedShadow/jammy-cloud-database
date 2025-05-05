package report

import (
	"context"
	"encoding/json"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/config"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	"go.uber.org/zap"
)

var DefaultInstance = NewInstance()

var HostName, _ = utils.Exec("hostname")

func SendMetricData(data *model.InputMetric) {
	topic := config.GetMonitorConfig().Topic + getVirtualMachineName()
	net := config.GetMonitorConfig().Net
	url := config.GetMonitorConfig().Url
	message, err := json.Marshal(data)
	if err != nil {
		logs.GetLogger().Error("json marshal error", zap.Error(err))
		return
	}
	err = DefaultInstance.Connect(context.Background(), net, url, topic)
	if err != nil {
		logs.GetLogger().Error("connect to kafka error", zap.Error(err))
		return
	}
	err = DefaultInstance.WriteMessage(message)
	if err != nil {
		logs.GetLogger().Error("write message to kafka failed", zap.Error(err), zap.String("kafka_address", url))
		return
	}
	_ = DefaultInstance.Close()
}

func getVirtualMachineName() string {
	return HostName
}
