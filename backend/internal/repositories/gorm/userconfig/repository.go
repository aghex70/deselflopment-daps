package userconfig

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

type Profile struct {
	Email string `gorm:"column:email"`
	UserConfig
}

type Tabler interface {
	TableName() string
}

func (UserConfig) TableName() string {
	return "daps_user_configs"
}

func (gr *UserConfigGormRepository) GetByUserId(ctx context.Context, userId int) (domain.Profile, error) {
	var p Profile
	query := fmt.Sprintf("SELECT daps_user_configs.auto_suggest, daps_user_configs.language, daps_users.email FROM daps_user_configs JOIN daps_users ON daps_user_configs.user_id = daps_users.id WHERE daps_users.id = %d", userId)

	result := gr.DB.Raw(query).Scan(&p)

	if result.Error != nil {
		return domain.Profile{}, result.Error
	}
	return p.ToDto(), nil
}

func (gr *UserConfigGormRepository) Update(ctx context.Context, uc domain.UserConfig, userId int) error {
	nuc := fromDto(uc)
	result := gr.DB.Model(&nuc).Where(UserConfig{UserId: userId}).Updates(map[string]interface{}{
		"auto_suggest": nuc.AutoSuggest,
		"language":     nuc.Language,
	})

	// We are always going to find the user config, so if the error is raised, it means that the user tried to update the configuration with the same existing values
	if result.RowsAffected == 0 {
		return errors.New("no changes were made")
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserConfigGormRepository) Create(ctx context.Context, uc domain.UserConfig) error {
	nuc := fromDto(uc)
	result := gr.DB.Create(&nuc)
	if result.Error != nil {
		return result.Error
	}
	return nil
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

func (p Profile) ToDto() domain.Profile {
	return domain.Profile{
		Email: p.Email,
		UserConfig: domain.UserConfig{
			ID:          p.ID,
			AutoSuggest: p.AutoSuggest,
			Language:    p.Language,
			UserId:      p.UserId,
		},
	}
}