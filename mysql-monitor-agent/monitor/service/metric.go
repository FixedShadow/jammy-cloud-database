package service

import (
	"fmt"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/collectors"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/config"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/model"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/report"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	"go.uber.org/zap"
	"time"
)

var collectorList = []collectors.CollectorInterface{
	&collectors.CPUCollector{},
	&collectors.MemCollector{},
	&collectors.DiskCollector{},
}

var HostName, _ = utils.Exec("hostname")

// CollectToServer start collect task
func CollectToServer() {
	ticker := time.NewTicker(time.Duration(config.DefaultMetricDeltaDataTimeInSecond) * time.Second)

	defer func() {
		if err := recover(); err != nil {
			logs.GetLogger().Panic("collect task panic", zap.String("reason", fmt.Sprintf("%v", err)))
		}
	}()
	for range ticker.C {
		go report.SendMetricData(collectMetricData())
	}
}

func collectMetricData() *model.InputMetric {
	var metricDatas []model.Metric
	now := utils.GetCurrTSInNano()
	for _, collector := range collectorList {
		collectedData := collector.Collect(now)
		if collectedData != nil {
			metricDatas = append(metricDatas, collectedData.Data...)
		}
	}
	data := &model.InputMetric{
		CollectTime: now,
		Data:        metricDatas,
		HostName:    HostName[:len(HostName)-1],
	}
	return data
}
