package config

import "github.com/spf13/viper"

type ConfigurationSettings struct {
	Port string
	Env  string
}

func New() ConfigurationSettings {
	viper.SetDefault("Port", "5000")
	viper.SetDefault("Env", "dev")
	return ConfigurationSettings{
		Port: viper.GetString("Port"),
		Env:  viper.GetString("Env"),
	}
}

func (c ConfigurationSettings) IsDev() bool {
	return c.Env == "dev"
}
