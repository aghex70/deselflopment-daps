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

func (gr *RelationshipGormRepository) CreateRelationships(ctx context.Context, userId int, categoryIds []int) error {
	var relationships []CategoryUserRelationship
	for _, categoryId := range categoryIds {
		relationships = append(relationships, CategoryUserRelationship{CategoryID: categoryId, UserID: userId})
	}
	result := gr.DB.Create(&relationships)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *RelationshipGormRepository) PurgeRelationships(ctx context.Context, userId int) error {
	result := gr.DB.Delete(&CategoryUserRelationship{}, "user_id = ?", userId)
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
