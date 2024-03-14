package config

import "github.com/spf13/viper"

type WorkerConfig struct {
	BrokerConfig CacheConfig
}

func LoadWorkerConfig() *WorkerConfig {
	cfg := &WorkerConfig{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil
	}
	return cfg
}
