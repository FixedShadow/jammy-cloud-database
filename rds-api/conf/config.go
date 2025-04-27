package conf

type ServerConfig struct {
	System    System    `yaml:"system"`
	LogConfig LogConfig `yaml:"log_config"`
}
