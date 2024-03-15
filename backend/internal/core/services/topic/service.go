package topic

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/topic"
	"log"
)

type Service struct {
	logger          *log.Logger
	topicRepository topic.Repository
}

func (s Service) Create(ctx context.Context, t domain.Topic) (domain.Topic, error) {
	t, err := s.topicRepository.Create(ctx, t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.Topic, error) {
	t, err := s.topicRepository.Get(ctx, id)
	if err != nil {
		return domain.Topic{}, err
	}
	return t, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	if err := s.topicRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Topic, error) {
	topics, err := s.topicRepository.List(ctx, filters)
	if err != nil {
		return []domain.Topic{}, err
	}
	return topics, nil
}

func (s Service) Update(ctx context.Context, id uint, fields *map[string]interface{}) error {
	if err := s.topicRepository.Update(ctx, id, fields); err != nil {
		return err
	}
	return nil
}

func NewTopicService(tr topic.Repository, logger *log.Logger) Service {
	return Service{
		logger:          logger,
		topicRepository: tr,
	}
}
