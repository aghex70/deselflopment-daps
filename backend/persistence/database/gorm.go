package database

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewGormDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Net, cfg.Host, cfg.Port, cfg.Name)
	log.Println(dsn)
	//dsn = "dapsuser:dapspw@tcp(db:11306)/dapsname"
	log.Println("Connecting ORM to database")
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting ORM to database %+v", err.Error())
		return nil, err
	}
	return gormDB, nil
}
