package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ConfigurationSettings struct {
	Port        string
	Env         string
	NewRelicKey string
}

func New() ConfigurationSettings {
	viper.SetDefault("Port", "5000")
	viper.SetDefault("Env", "development")
	viper.SetDefault("NewRelicKey", "6b74b13f1609b362e7129c9081e6e40a9ee6NRAL")

	return ConfigurationSettings{
		Port:        viper.GetString("Port"),
		Env:         viper.GetString("Env"),
		NewRelicKey: viper.GetString("NewRelicKey"),
	}
}

func (c ConfigurationSettings) IsDev() bool {
	return c.Env == "development"
}

func (c ConfigurationSettings) DisplaySettings() {
	fmt.Println("Application Settings:")
	fmt.Println("⇨ " + c.Env)
	fmt.Println("⇨ " + c.Port)
}
