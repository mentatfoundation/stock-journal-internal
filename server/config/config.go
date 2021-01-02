package config

import "github.com/spf13/viper"

type ConfigurationSettings struct {
	Port         string
	Env          string
	RollbarToken string
	SentryDsn    string
}

func New() ConfigurationSettings {
	viper.SetDefault("Port", "5000")
	viper.SetDefault("Env", "qa")
	viper.SetDefault("RollbarToken", "580a34ca2b52410c8df71318ac58e2f5")
	viper.SetDefault("SentryDsn", "https://06b83a9a0fd045bbaf6b503fdefc8a66@o499035.ingest.sentry.io/5577124")

	return ConfigurationSettings{
		Port:         viper.GetString("Port"),
		Env:          viper.GetString("Env"),
		RollbarToken: viper.GetString("RollbarToken"),
		SentryDsn:    viper.GetString("SentryDsn"),
	}
}

func (c ConfigurationSettings) IsDev() bool {
	return c.Env == "development"
}
