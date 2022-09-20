package config

import (
	"github.com/spf13/viper"
	"log"
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
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(CONFIG_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%+v", err.Error())
	}

	return &Config{
		Cache:     LoadCacheConfig(),
		Database:  LoadDatabaseConfig(),
		Logger:    LoadLoggerConfig(),
		Providers: LoadProvidersConfig(),
		Server:    LoadServerConfig(),
	}, nil
}
