package userconfig

import (
	"context"
	"log"
	"net/http"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	uc "github.com/aghex70/daps/internal/repositories/gorm/userconfig"
	"github.com/aghex70/daps/server"
)

type Service struct {
	logger               *log.Logger
	userConfigRepository *uc.GormRepository
}

func (s Service) Update(ctx context.Context, r *http.Request, req ports.UpdateUserConfigRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	uConfig := domain.UserConfig{
		UserId:      int(userId),
		AutoSuggest: req.AutoSuggest,
		AutoRemind:  req.AutoRemind,
		Language:    req.Language,
	}
	err := s.userConfigRepository.Update(ctx, uConfig, int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request) (domain.Profile, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	p, err := s.userConfigRepository.GetByUserId(ctx, int(userId))
	if err != nil {
		return domain.Profile{}, err
	}
	return p, nil
}

func NewUserConfigService(ucr *uc.GormRepository, logger *log.Logger) Service {
	return Service{
		logger:               logger,
		userConfigRepository: ucr,
	}
}
