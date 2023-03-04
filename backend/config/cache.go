package config

import (
	"github.com/spf13/viper"
)

type CacheConfig struct {
	DB       int    `mapstructure:"CACHE_DB"`
	Host     string `mapstructure:"CACHE_HOST"`
	Port     int    `mapstructure:"CACHE_PORT"`
	Name     string `mapstructure:"CACHE_NAME"`
	User     string `mapstructure:"CACHE_USER"`
	Password string `mapstructure:"CACHE_PASSWORD"`
}

func LoadCacheConfig() *CacheConfig {
	cfg := &CacheConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
