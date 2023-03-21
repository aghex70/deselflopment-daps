package config

import "github.com/spf13/viper"

type RestConfig struct {
	Host string `mapstructure:"BACKEND_HOST"`
	Port int    `mapstructure:"BACKEND_PORT"`
}

func LoadRestConfig() *RestConfig {
	cfg := &RestConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	return cfg
}
