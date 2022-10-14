package todo

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

func (s TodoService) CheckExistentTodo(ctx context.Context, name string, categoryId int) (domain.Todo, bool) {
	t, err := s.todoRepository.GetByNameAndCategory(ctx, name, categoryId)
	return t, !errors.Is(err, gorm.ErrRecordNotFound)
}
