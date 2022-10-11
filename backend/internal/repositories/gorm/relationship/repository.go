package relationship

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"log"
)

type RelationshipGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type CategoryUserRelationship struct {
	ID         int `gorm:"primaryKey;column:id"`
	CategoryID int `gorm:"column:category_id"`
	UserID     int `gorm:"column:user_id"`
}

type CategoryUserRelationship2 struct {
	CategoryID int
	UserID     int
}

type Tabler interface {
	TableName() string
}

func (CategoryUserRelationship) TableName() string {
	return "daps_categories_users_relationships"
}

func (gr *RelationshipGormRepository) CreateRelationships(ctx context.Context, userId int) error {
	var relationships = []CategoryUserRelationship{{CategoryID: 1, UserID: userId}, {CategoryID: 2, UserID: userId}, {CategoryID: 3, UserID: userId}, {CategoryID: 4, UserID: userId}, {CategoryID: 5, UserID: userId}}
	result := gr.DB.Create(&relationships)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewRelationshipGormRepository(db *gorm.DB) (*RelationshipGormRepository, error) {
	return &RelationshipGormRepository{
		DB: db,
	}, nil
}
