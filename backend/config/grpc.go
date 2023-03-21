package config

import "github.com/spf13/viper"

type GrpcConfig struct {
}

func LoadGrpcConfig() *GrpcConfig {
	cfg := &GrpcConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil
	}
	return cfg
}
