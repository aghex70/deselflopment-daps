package category

import (
	"context"
	"errors"
)

func (s CategoryService) ValidateCreation(ctx context.Context, name string, userId int) error {
	cc, err := s.categoryRepository.ListCustom(ctx, userId)
	if err != nil {
		return err
	}
	if len(cc) > 2 {
		return errors.New("max number of custom categories reached")
	}

	if _, err := s.categoryRepository.GetBaseCategory(ctx, name); err == nil {
		return errors.New("already existent base category with that name")
	}
	return nil
}
