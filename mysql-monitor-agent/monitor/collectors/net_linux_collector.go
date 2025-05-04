package collectors

type NetStates struct {
	byteSent        float64
	byteRecv        float64
	packetSent      float64
	packetRecv      float64
	errIn           float64
	errOut          float64
	dropIn          float64
	dropOut         float64
	collectTime     int64
	uptimeInSeconds int64
}

type NetCollector struct {
	LastStates *NetStates
}
