package userconfig

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
)

type UserConfigGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type UserConfig struct {
	ID          int    `gorm:"primaryKey;column:id"`
	UserId      int    `gorm:"column:user_id"`
	AutoSuggest bool   `gorm:"column:auto_suggest"`
	Language    string `gorm:"column:language"`
}

type Tabler interface {
	TableName() string
}

func (UserConfig) TableName() string {
	return "daps_user_configs"
}

func (gr *UserConfigGormRepository) GetByUserId(ctx context.Context, userId int) (domain.UserConfig, error) {
	var uc UserConfig
	result := gr.DB.Where(&UserConfig{UserId: userId}).First(&uc)
	if result.Error != nil {
		return domain.UserConfig{}, result.Error
	}

	return uc.ToDto(), nil
}

func (gr *UserConfigGormRepository) Update(ctx context.Context, uc domain.UserConfig, userId int) error {
	nuc := fromDto(uc)
	result := gr.DB.Model(&nuc).Where(UserConfig{UserId: userId}).Updates(map[string]interface{}{
		"auto_suggest": nuc.AutoSuggest,
		"language":     nuc.Language,
	})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserConfigGormRepository) Create(ctx context.Context, uc domain.UserConfig) (domain.UserConfig, error) {
	nuc := fromDto(uc)
	result := gr.DB.Create(&nuc)
	if result.Error != nil {
		return domain.UserConfig{}, result.Error
	}
	return nuc.ToDto(), nil
}

func NewUserConfigGormRepository(db *gorm.DB) (*UserConfigGormRepository, error) {
	return &UserConfigGormRepository{
		DB: db,
	}, nil
}

func (uc UserConfig) ToDto() domain.UserConfig {
	return domain.UserConfig{
		ID:          uc.ID,
		AutoSuggest: uc.AutoSuggest,
		Language:    uc.Language,
		UserId:      uc.UserId,
	}
}

func fromDto(uc domain.UserConfig) UserConfig {
	return UserConfig{
		ID:          uc.ID,
		AutoSuggest: uc.AutoSuggest,
		Language:    uc.Language,
		UserId:      uc.UserId,
	}
}
