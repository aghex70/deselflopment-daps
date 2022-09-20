package config

import "github.com/spf13/viper"

type GrpcConfig struct {
}

func LoadGrpcConfig() *GrpcConfig {
	cfg := &GrpcConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
