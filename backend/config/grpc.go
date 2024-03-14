package config

import "github.com/spf13/viper"

type GrpcConfig struct {
}

func LoadGrpcConfig() *GrpcConfig {
	cfg := &GrpcConfig{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil
	}
	return cfg
}
