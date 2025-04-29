package conf

type DubboConfig struct {
	RegistryAddress string `yaml:"registry_address"`
	ClientName      string `yaml:"client_name"`
}
