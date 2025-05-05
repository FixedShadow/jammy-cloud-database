package collectors

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/model"
	"github.com/shirou/gopsutil/v3/mem"
	"go.uber.org/zap"
)

type MemCollector struct {
}

func (m *MemCollector) Collect(collectTime int64) *model.InputMetric {
	memory, err := mem.VirtualMemory()
	if err != nil {
		logs.GetLogger().Error("Get memory stats failed", zap.Error(err))
		return nil
	}
	metricsDatas := []model.Metric{
		{
			MetricName:  "MemoryTotalSpace",
			MetricValue: float64(memory.Total),
		},
		{
			MetricName:  "MemoryFreeSpace",
			MetricValue: float64(memory.Available),
		},
		{
			MetricName:  "MemoryUsedSpace",
			MetricValue: float64(memory.Total - memory.Available),
		},
		{
			MetricName:  "MemoryBuffers",
			MetricValue: float64(memory.Buffers),
		},
		{
			MetricName:  "MemoryCached",
			MetricValue: float64(memory.Cached),
		},
		{
			MetricName:  "MemoryUsedUtilization",
			MetricValue: float64(memory.Total-memory.Available) / float64(memory.Total) * model.ToPercent,
		},
		{
			MetricName:  "MemoryFreeUtilization",
			MetricValue: float64(memory.Available) / float64(memory.Total) * model.ToPercent,
		},
	}
	return &model.InputMetric{
		Data:        metricsDatas,
		Type:        "memory",
		CollectTime: collectTime,
	}
}
