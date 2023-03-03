package config

import "github.com/spf13/viper"

type WorkerConfig struct {
	BrokerConfig CacheConfig
}

func LoadWorkerConfig() *WorkerConfig {
	cfg := &WorkerConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
