package collectors

import "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/model"

type CollectorInterface interface {
	Collect(collectTime int64) *model.InputMetric
}
