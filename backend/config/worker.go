package config

import "github.com/spf13/viper"

type WorkerConfig struct {
	BrokerConfig CacheConfig
}

func LoadWorkerConfig() *WorkerConfig {
	cfg := &WorkerConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	if err != nil {
		return nil
	}
	return cfg
}
