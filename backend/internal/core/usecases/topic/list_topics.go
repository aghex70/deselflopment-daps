package topic

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"

	//"github.com/aghex70/daps/server"
	"log"
)

type ListTopicsUseCase struct {
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *ListTopicsUseCase) Execute(ctx context.Context, filters *map[string]interface{}, userID uint) ([]domain.Topic, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Topic{}, err
	}

	if !u.Active {
		return []domain.Topic{}, pkg.InactiveUserError
	}

	// Set the user ID into the fields map (retrieve only own topics)
	(*filters)["owner_id"] = userID

	topics, err := uc.TopicService.List(ctx, filters)
	if err != nil {
		return []domain.Topic{}, err
	}
	return topics, nil
}

func NewListTopicsUseCase(s topic.Servicer, u user.Servicer, logger *log.Logger) *ListTopicsUseCase {
	return &ListTopicsUseCase{
		TopicService: s,
		UserService:  u,
		logger:       logger,
	}
}
