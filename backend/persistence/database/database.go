package database

import (
	"database/sql"
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/go-sql-driver/mysql"
	"log"
)

func NewSqlDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	mySqlConfig := mysql.NewConfig()
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	mySqlConfig.Addr = address
	mySqlConfig.DBName = cfg.Name
	mySqlConfig.User = cfg.User
	mySqlConfig.Passwd = cfg.Password
	mySqlConfig.Net = cfg.Net
	mySqlConfig.ParseTime = true

	log.Println("Connecting to database")
	sqlDB, err := sql.Open(cfg.Dialect, mySqlConfig.FormatDSN())
	if err != nil {
		log.Fatalf("Error connecting to database %+v", err.Error())
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifeTime)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	return sqlDB, nil
}
