package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

func (s Service) CheckExistentTodo(ctx context.Context, name string, categoryID int) (domain.Todo, bool) {
	//t, err := s.todoRepository.GetByNameAndCategory(ctx, name, categoryID)
	//return t, !errors.Is(err, gorm.ErrRecordNotFound)
	return domain.Todo{}, false
}

func (s Service) CheckCategoryPermissions(ctx context.Context, userID, categoryID int) error {
	//err := s.relationshipRepository.GetUserCategory(ctx, userID, categoryID)
	//return err
	return nil
}

func (s Service) CheckCategoriesPermissions(ctx context.Context, userID int) ([]int, error) {
	//categoryIDs, err := s.relationshipRepository.ListUserCategories(ctx, userID)
	//return categoryIDs, err
	return nil, nil
}
