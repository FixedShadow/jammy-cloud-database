package model

type ContainerCreateSpecs struct {
	ContainerName string `json:"containerName"`
	ContainerType string `json:"containerType"`
	CpuNum        int    `json:"cpuNum"`
	Memory        int    `json:"memory"`
	DiskSize      int    `json:"diskSize"`
	ImageId       string `json:"imageId"`
	ImageType     string `json:"imageType"`
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
