package category

import (
	"database/sql"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type GormRepository struct {
	*gorm.DB
	SqlDb *sql.DB
}

func NewGormRepository(db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
		DB: db,
	}, nil
}
