package collectors

import "sync"

type DiskCollector struct {
	DiskMap sync.Map
}

type DiskIOCountersStat struct {
	collectTime     int64
	uptimeInSeconds int64
	readBytes       float64
	readCount       float64
	writeBytes      float64
	writeCount      float64
}
