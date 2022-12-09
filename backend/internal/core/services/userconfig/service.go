package userconfig

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	uc "github.com/aghex70/daps/internal/repositories/gorm/userconfig"
	"github.com/aghex70/daps/server"
	"log"
	"net/http"
)

type UserConfigService struct {
	logger               *log.Logger
	userConfigRepository *uc.UserConfigGormRepository
}

func (s UserConfigService) Update(ctx context.Context, r *http.Request, req ports.UpdateUserConfigRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	uc := domain.UserConfig{
		UserId:      int(userId),
		AutoSuggest: req.AutoSuggest,
		Language:    req.Language,
	}
	err := s.userConfigRepository.Update(ctx, uc, int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s UserConfigService) Get(ctx context.Context, r *http.Request, req ports.GetUserConfigRequest) (domain.UserConfig, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	td, err := s.userConfigRepository.GetByUserId(ctx, int(userId))
	if err != nil {
		return domain.UserConfig{}, err
	}
	return td, nil
}

func NewUserConfigService(ucr *uc.UserConfigGormRepository, logger *log.Logger) UserConfigService {
	return UserConfigService{
		logger:               logger,
		userConfigRepository: ucr,
	}
}
