package conf

type Registry struct {
	ServiceName string `yaml:"service_name"`
	Type        string `yaml:"type"`
	Address     string `yaml:"address"`
	Port        string `yaml:"port"`
}
