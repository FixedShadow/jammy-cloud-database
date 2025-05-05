package collectors

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/model"
	"github.com/shirou/gopsutil/v3/load"
	"go.uber.org/zap"
	"runtime"
)

type LoadCollector struct {
}

func (l *LoadCollector) Collect(collectTime int64) *model.InputMetric {
	loadAvg, err := load.Avg()
	if nil != err {
		logs.GetLogger().Info("get system load error", zap.Error(err))
		return nil
	}

	numCPU := float64(runtime.NumCPU())

	metricsDatas := []model.Metric{
		{MetricName: "Load1m", MetricValue: loadAvg.Load1},
		{MetricName: "Load5m", MetricValue: loadAvg.Load5},
		{MetricName: "Load15m", MetricValue: loadAvg.Load15},
		{MetricName: "LoadPerCore1m", MetricValue: loadAvg.Load1 / numCPU},
		{MetricName: "LoadPerCore5m", MetricValue: loadAvg.Load5 / numCPU},
		{MetricName: "LoadPerCore15m", MetricValue: loadAvg.Load15 / numCPU},
	}

	return &model.InputMetric{
		Data:        metricsDatas,
		Type:        "load",
		CollectTime: collectTime,
	}
}
