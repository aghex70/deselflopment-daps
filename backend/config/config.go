package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	CONFIG_NAME string = ".env"
	CONFIG_TYPE string = "env"
	CONFIG_PATH string = "./"
)

type Config struct {
	Cache     *CacheConfig
	Database  *DatabaseConfig
	Logger    *LoggerConfig
	Providers *ProvidersConfig
	Server    *ServerConfig
	Broker    *CacheConfig
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%+v", err.Error())
	}

	return &Config{
		Cache:     LoadCacheConfig(),
		Database:  LoadDatabaseConfig(),
		Logger:    LoadLoggerConfig(),
		Providers: LoadProvidersConfig(),
		Server:    LoadServerConfig(),
		Broker:    LoadCacheConfig(),
	}, nil
}
