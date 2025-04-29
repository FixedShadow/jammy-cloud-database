package conf

type ServerConfig struct {
	System    System    `yaml:"system"`
	Registry  Registry  `yaml:"registry"`
	LogConfig LogConfig `yaml:"log_config"`
}
