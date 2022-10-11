package category

import (
	"context"
	"errors"
)

func (s CategoryService) ValidateCreation(ctx context.Context, name string, userId int) error {
	if _, err := s.categoryRepository.GetUserCategory(ctx, name, userId); err == nil {
		return errors.New("already existent category with that user and name")
	}
	return nil
}

func (s CategoryService) ValidateModification(ctx context.Context, name string, userId int) error {
	if _, err := s.categoryRepository.GetBaseCategory(ctx, name); err == nil {
		return errors.New("already existent base category with that name")
	}
	return nil
}
