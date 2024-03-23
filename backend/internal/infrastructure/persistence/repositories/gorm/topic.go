package gorm

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Topic struct {
	gorm.Model
	ID      uint
	Name    string
	OwnerID uint
}

func (t Topic) ToDto() domain.Topic {
	var createdAt time.Time
	if !t.CreatedAt.IsZero() {
		createdAt = t.CreatedAt
	}

	return domain.Topic{
		ID:        t.ID,
		CreatedAt: createdAt,
		Name:      t.Name,
		OwnerID:   t.OwnerID,
	}
}

func TopicFromDto(c domain.Topic) Topic {
	return Topic{
		ID:      c.ID,
		Name:    c.Name,
		OwnerID: c.OwnerID,
	}
}

func (Topic) TableName() string {
	return "daps_topics"
}

func (gr *TopicRepository) Create(ctx context.Context, t domain.Topic) (domain.Topic, error) {
	nt := TopicFromDto(t)
	if result := gr.DB.Create(&nt); result.Error != nil {
		return domain.Topic{}, result.Error
	}

	return nt.ToDto(), nil
}

func (gr *TopicRepository) Get(ctx context.Context, id uint) (domain.Topic, error) {
	var t Topic
	if result := gr.DB.First(&t, id); result.Error != nil {
		return domain.Topic{}, result.Error
	}
	return t.ToDto(), nil
}

func (gr *TopicRepository) Delete(ctx context.Context, id uint) error {
	// Fetch the topic along with its associations
	if err := gr.DB.Exec(
		"DELETE FROM daps_note_topics WHERE topic_id = ?", id).Error; err != nil {
		return err
	}

	var topic Topic
	if result := gr.DB.Delete(&topic, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *TopicRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Topic, error) {
	var ts []Topic
	var topics []domain.Topic

	query := gr.DB
	if filters != nil {
		// Convert map[string]interface{} to a slice of arguments
		var args []interface{}
		var conditions []string
		for key, value := range *filters {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}
		query = query.Where(strings.Join(conditions, " AND "), args...)
	}

	if result := query.Find(&ts); result.Error != nil {
		return []domain.Topic{}, result.Error
	}

	for _, t := range ts {
		ts := t.ToDto()
		topics = append(topics, ts)
	}
	return topics, nil
}

func (gr *TopicRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var t Topic
	t.ID = id
	if result := gr.DB.Model(&t).Updates(*filters); result.Error != nil {
		return result.Error
	}
	return nil
}

type TopicRepository struct {
	*gorm.DB
}

func NewGormTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{db}
}
