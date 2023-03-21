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

func (s TodoService) CheckCategoryPermissions(ctx context.Context, userId, categoryId int) error {
	err := s.relationshipRepository.GetUserCategory(ctx, userId, categoryId)
	return err
}

func (s TodoService) CheckCategoriesPermissions(ctx context.Context, userId int) ([]int, error) {
	categoryIds, err := s.relationshipRepository.ListUserCategories(ctx, userId)
	return categoryIds, err
}
