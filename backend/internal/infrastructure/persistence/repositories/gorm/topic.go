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
	Name    string
	OwnerID uint
}

func (c Topic) ToDto() domain.Topic {
	var createdAt time.Time
	if !c.CreatedAt.IsZero() {
		createdAt = c.CreatedAt
	}

	return domain.Topic{
		ID:        c.ID,
		CreatedAt: createdAt,
		Name:      c.Name,
		OwnerID:   c.OwnerID,
	}
}

func TopicFromDto(c domain.Topic) Topic {
	return Topic{
		Name:    c.Name,
		OwnerID: c.OwnerID,
	}
}

func (Topic) TableName() string {
	return "daps_topics"
}

func (gr *TopicRepository) Create(ctx context.Context, c domain.Topic) (domain.Topic, error) {
	nc := TopicFromDto(c)
	if result := gr.DB.Create(&nc); result.Error != nil {
		return domain.Topic{}, result.Error
	}

	return nc.ToDto(), nil
}

func (gr *TopicRepository) Get(ctx context.Context, id uint) (domain.Topic, error) {
	var c Topic
	if result := gr.DB.First(&c, id); result.Error != nil {
		return domain.Topic{}, result.Error
	}

	return c.ToDto(), nil
}

func (gr *TopicRepository) Delete(ctx context.Context, id uint) error {
	// Fetch the topic along with its associations
	if err := gr.DB.Exec(
		"DELETE FROM daps_topic_users WHERE topic_id = ?", id).Error; err != nil {
		return err
	}

	var topic Topic
	if result := gr.DB.Delete(&topic, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *TopicRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Topic, error) {
	var cs []Topic
	var cats []domain.Topic

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

	if result := query.Find(&cs); result.Error != nil {
		return []domain.Topic{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *TopicRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var c Topic
	c.ID = id
	if result := gr.DB.Model(&c).Updates(*filters); result.Error != nil {
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
