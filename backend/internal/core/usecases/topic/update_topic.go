package topic

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/topic"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/topic"

	"log"
)

type UpdateTopicUseCase struct {
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *UpdateTopicUseCase) Execute(ctx context.Context, r requests.UpdateTopicRequest, userID uint) error {
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
	owner := utils.IsTopicOwner(t.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	fields := common.StructToMap(r, "topic_id")
	if err = uc.TopicService.Update(ctx, t.ID, &fields); err != nil {
		return err
	}
	return nil
}

func NewUpdateTopicUseCase(s topic.Servicer, u user.Servicer, logger *log.Logger) *UpdateTopicUseCase {
	return &UpdateTopicUseCase{
		TopicService: s,
		UserService:  u,
		logger:       logger,
	}
}
