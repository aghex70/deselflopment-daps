package topic

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/topic"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/topic"

	"log"
)

type DeleteTopicUseCase struct {
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *DeleteTopicUseCase) Execute(ctx context.Context, r requests.DeleteTopicRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	t, err := uc.TopicService.Get(ctx, r.TopicID)
	if err != nil {
		return err
	}
	owner := utils.IsTopicOwner(t.OwnerID, u.ID)
	if !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.TopicService.Delete(ctx, r.TopicID); err != nil {
		return err
	}
	return nil
}

func NewDeleteTopicUseCase(s topic.Servicer, u user.Servicer, logger *log.Logger) *DeleteTopicUseCase {
	return &DeleteTopicUseCase{
		TopicService: s,
		UserService:  u,
		logger:       logger,
	}
}
