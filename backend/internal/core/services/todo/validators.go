package todo

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

func (s TodoService) CheckExistentTodo(ctx context.Context, name string, link string, userId int) (domain.Todo, bool) {
	t, err := s.todoRepository.GetByNameAndLink(ctx, name, link, userId)
	return t, !errors.Is(err, gorm.ErrRecordNotFound)
}
