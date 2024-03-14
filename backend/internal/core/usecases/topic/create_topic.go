package topic

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/topic"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type CreateTopicUseCase struct {
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *CreateTopicUseCase) Execute(ctx context.Context, userID uint, r requests.CreateTopicRequest) (domain.Topic, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Topic{}, err
	}

	if !u.Active {
		return domain.Topic{}, pkg.InactiveUserError
	}

	t := domain.Topic{
		Name:    r.Name,
		OwnerID: u.ID,
	}
	t, err = uc.TopicService.Create(ctx, t)
	if err != nil {
		return domain.Topic{}, err
	}

	return t, nil
}

func NewCreateTopicUseCase(s topic.Servicer, u user.Servicer, logger *log.Logger) *CreateTopicUseCase {
	return &CreateTopicUseCase{
		TopicService: s,
		UserService:  u,
		logger:       logger,
	}
}
