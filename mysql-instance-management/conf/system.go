package conf

type System struct {
	BindAddress string `yaml:"bind_address"`
	Port        int    `yaml:"port"`
}
