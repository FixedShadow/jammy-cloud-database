package dto

type ContainerSpec struct {
	Name    string `json:"name"` //容器名称 使用uuid生成(删除“-”)
	Type    string `json:"type"`
	CpuNum  int    `json:"cpuNum"`
	Memory  int    `json:"memory"`
	Size    int64  `json:"size"`
	ImageId string `json:"imageId"`
}

type ContainerOperate struct {
	Timeout int    `json:"timeout"`
	Action  string `json:"action"`
}
