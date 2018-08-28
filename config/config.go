package config

import "github.com/spf13/viper"

var Config *config

func init() {
	Config = &config{}
}

type config struct{}

func (c *config) BuildVersion() string {
	return viper.GetString("build.version")
}

func (c *config) BuildRevision() string {
	return viper.GetString("build.revision")
}

func (c *config) BuildTime() string {
	return viper.GetString("build.time")
}
