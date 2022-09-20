package config

import (
	"github.com/spf13/viper"
	"time"
)

type DatabaseConfig struct {
	Dialect            string `mapstructure:"DB_DIALECT"`
	Host               string `mapstructure:"DB_HOST"`
	LogQuery           bool
	MaxOpenConnections int           `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConnections int           `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	MaxConnLifeTime    time.Duration `mapstructure:"DB_MAX_CONN_LIFE_TIME"`
	MigrationDir       string        `mapstructure:"DB_MIGRATION_DIR"`
	Name               string        `mapstructure:"DB_NAME"`
	Net                string        `mapstructure:"DB_NETWORK"`
	Port               int           `mapstructure:"DB_PORT"`
	Password           string        `mapstructure:"DB_PASSWORD"`
	User               string        `mapstructure:"DB_USER"`
}

func LoadDatabaseConfig() *DatabaseConfig {
	cfg := &DatabaseConfig{}
	viper.Unmarshal(cfg)
	return cfg
}
