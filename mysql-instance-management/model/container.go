package model

type ContainerSpecs struct {
	ContainerName string `json:"containerName"`
	ContainerType string `json:"containerType"`
	CpuNum        int    `json:"cpuNum"`
	Memory        int    `json:"memory"`
	Disk          int    `json:"disk"`
	ImageId       string `json:"imageId"`
	Type          string `json:"type"`
}

type ContainerInfo struct {
	ContainerName  string `json:"containerName"`
	ManageNic      string `json:"manageNic"`
	DataNic        string `json:"dataNic"`
	Status         string `json:"status"`
	NetworkStorage bool   `json:"networkStorage"`
	CreateAt       int64  `json:"createAt"`
}
