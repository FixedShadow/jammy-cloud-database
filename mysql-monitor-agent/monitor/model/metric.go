package model

const (
	// ByteToBit shift Byte to Bit
	ByteToBit = 8
	// KBToByte shift KB to Bit
	KBToByte = 1024
	// MBToByte shift MB to Byte
	MBToByte = 1024 * 1024
	// GBToByte shift GB to Byte
	GBToByte = 1024 * 1024 * 1024
	// ToPercent shift to %
	ToPercent = 100
	// WattsTomW Watts To Milli watts
	WattsTomW = 1000
)

type Metric struct {
	MetricName   string  `json:"metric_name"`
	MetricValue  float64 `json:"metric_value"`
	MetricPrefix string  `json:"metric_prefix,omitempty"`
}

type InputMetric struct {
	CollectTime int64    `json:"collect_time"`
	Type        string   `json:"-"`
	Data        []Metric `json:"data"`
}
