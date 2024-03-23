package gorm

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Note struct {
	gorm.Model
	Content string
	Users   []User `gorm:"many2many:daps_note_users;save_association:true"`
	OwnerID uint
	Topics  []Topic `gorm:"many2many:daps_note_topics;save_association:true"`
	Shared  bool
}

func (n Note) ToDto() domain.Note {
	var createdAt time.Time
	if !n.CreatedAt.IsZero() {
		createdAt = n.CreatedAt
	}

	var topics []domain.Topic
	if n.Topics != nil {
		for _, topic := range n.Topics {
			topics = append(topics, topic.ToDto())
		}
	}

	var users []domain.User
	if n.Users != nil {
		for _, user := range n.Users {
			users = append(users, user.ToDto())
		}
	}

	return domain.Note{
		ID:        n.ID,
		CreatedAt: createdAt,
		Content:   n.Content,
		OwnerID:   n.OwnerID,
		Users:     &users,
		Topics:    topics,
		Shared:    n.Shared,
	}
}

func NoteFromDto(c domain.Note) Note {
	var users []User
	if c.Users != nil {
		for _, userDTO := range *c.Users {
			user := UserFromDto(userDTO)
			users = append(users, user)
		}
	}

	var topics []Topic
	if c.Topics != nil {
		for _, topicDTO := range c.Topics {
			topic := TopicFromDto(topicDTO)
			topics = append(topics, topic)
		}
	}

	return Note{
		Content: c.Content,
		OwnerID: c.OwnerID,
		Users:   users,
		Topics:  topics,
		Shared:  c.Shared,
	}
}

func (Note) TableName() string {
	return "daps_notes"
}

func (gr *NoteRepository) Create(ctx context.Context, n domain.Note) (domain.Note, error) {
	nc := NoteFromDto(n)
	if result := gr.DB.Create(&nc); result.Error != nil {
		return domain.Note{}, result.Error
	}

	// Hack to get around the fact that GORM doesn't support many-to-many relationships
	if nc.Users == nil {
		if err := gr.DB.Association("Users").Append(nc.Users); err != nil {

			return domain.Note{}, err
		}
	}

	return nc.ToDto(), nil
}

func (gr *NoteRepository) Get(ctx context.Context, id uint) (domain.Note, error) {
	var n Note
	if result := gr.DB.First(&n, id); result.Error != nil {
		return domain.Note{}, result.Error
	}

	// Retrieve users associated with the note if they exist
	if n.Users == nil {
		if err := gr.DB.Model(&n).Association("Users").Find(&n.Users); err != nil {
			return domain.Note{}, err
		}
	}

	// Retrieve topics associated with the note if they exist
	if n.Topics == nil {
		if err := gr.DB.Model(&n).Association("Topics").Find(&n.Topics); err != nil {
			return domain.Note{}, err
		}
	}

	return n.ToDto(), nil
}

func (gr *NoteRepository) Delete(ctx context.Context, id uint) error {
	// Fetch the note along with its associations
	if err := gr.DB.Exec(
		"DELETE FROM daps_note_users WHERE note_id = ?", id).Error; err != nil {
		return err
	}

	var note Note
	if result := gr.DB.Delete(&note, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *NoteRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Note, error) {
	var cs []Note
	var cats []domain.Note

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

	// Preload topics
	query = query.Preload("Topics")

	if result := query.Find(&cs); result.Error != nil {
		return []domain.Note{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *NoteRepository) Update(ctx context.Context, id uint, ud domain.Note) error {
	var n Note
	if err := gr.DB.First(&n, id).Error; err != nil {
		return err
	}

	n.Content = ud.Content
	n.Topics = make([]Topic, 0)

	// Clean manually the many-to-many relationship
	if err := gr.DB.Exec("DELETE FROM daps_note_topics WHERE note_id = ?", id).Error; err != nil {
		return err
	}

	for _, topic := range ud.Topics {
		t := TopicFromDto(topic)
		n.Topics = append(n.Topics, t)
	}
	n.OwnerID = ud.OwnerID
	n.Shared = ud.Shared
	n.Users = make([]User, 0)

	// Clean manually the many-to-many relationship
	if err := gr.DB.Exec("DELETE FROM daps_note_users WHERE note_id = ?", id).Error; err != nil {
		return err
	}

	for _, user := range *ud.Users {
		u := UserFromDto(user)
		n.Users = append(n.Users, u)
	}

	if result := gr.DB.Save(&n); result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *NoteRepository) Share(ctx context.Context, id uint, u domain.User) error {
	var n Note
	if result := gr.DB.Preload("Users").First(&n, id); result.Error != nil {
		return result.Error
	}

	if n.Users == nil {
		n.Users = make([]User, 0)
	}

	// Check if the user is already in the note
	if !isUserInNote(n.Users, u) {
		query := fmt.Sprintf("INSERT INTO daps_note_users (note_id, user_id) VALUES (%d, %d)", id, u.ID)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
		// Update the note to be shared
		query = fmt.Sprintf("UPDATE daps_notes SET shared = TRUE WHERE id = %d", id)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
	}
	return nil
}

func isUserInNote(users []User, u domain.User) bool {
	for _, user := range users {
		if user.ID == u.ID {
			return true
		}
	}
	return false
}

func (gr *NoteRepository) Unshare(ctx context.Context, id uint, u domain.User) error {
	if err := gr.DB.Exec(
		"DELETE FROM daps_note_users WHERE user_id = ? AND note_id = ?", u.ID, id).Error; err != nil {
		return err
	}

	var n Note
	if result := gr.DB.Preload("Users").First(&n, id); result.Error != nil {
		return result.Error
	}

	if n.Users == nil || len(n.Users) == 1 {
		query := fmt.Sprintf("UPDATE daps_notes SET shared = FALSE WHERE id = %d", id)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}

type NoteRepository struct {
	*gorm.DB
}

func NewGormNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db}
}
