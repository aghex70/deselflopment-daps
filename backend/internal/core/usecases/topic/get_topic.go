package topic

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/topic"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/topic"

	"log"
)

type GetTopicUseCase struct {
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *GetTopicUseCase) Execute(ctx context.Context, r requests.GetTopicRequest, userID uint) (domain.Topic, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Topic{}, err
	}

	if !u.Active {
		return domain.Topic{}, pkg.InactiveUserError
	}

	t, err := uc.TopicService.Get(ctx, r.TopicID)
	if err != nil {
		return domain.Topic{}, err
	}
	if owner := utils.IsTopicOwner(t.OwnerID, u.ID); !owner {
		return domain.Topic{}, pkg.UnauthorizedError
	}

	return t, nil
}

func NewGetTopicUseCase(s topic.Servicer, u user.Servicer, logger *log.Logger) *GetTopicUseCase {
	return &GetTopicUseCase{
		TopicService: s,
		UserService:  u,
		logger:       logger,
	}
}
