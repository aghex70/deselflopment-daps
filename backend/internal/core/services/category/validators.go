package category

import (
	"context"
	"errors"
	"fmt"
)

func (s CategoryService) ValidateCreation(ctx context.Context, name string, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_categories.name = '%s'", userId, name)
	fmt.Println(conditions)
	categoryId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	if err != nil {
		return err
	}
	if categoryId != 0 {
		return errors.New("already existent category with that user and name")
	}
	return nil
}

func (s CategoryService) ValidateModification(ctx context.Context, categoryId, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d AND daps_categories.owner_id = %d", userId, categoryId, userId)
	catId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	if err != nil {
		return err
	}
	if catId == 0 {
		return errors.New("cannot update category")
	}
	return nil
}

func (s CategoryService) ValidateShare(ctx context.Context, categoryId, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d AND daps_categories.owner_id = %d", userId, categoryId, userId)
	catId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	if err != nil {
		return err
	}
	if catId == 0 {
		return errors.New("cannot update category")
	}
	return nil
}

func (s CategoryService) ValidateUnshare(ctx context.Context, categoryId, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d", userId, categoryId)
	catId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	if err != nil {
		return err
	}
	if catId == 0 {
		return errors.New("cannot update category")
	}
	return nil
}

func (s CategoryService) ValidateRetrieval(ctx context.Context, categoryId, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d", userId, categoryId)
	catId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	fmt.Println("rrrrrrrrrrrrrrrrrrrrrrrr")
	if err != nil {
		return err
	}
	if catId == 0 {
		fmt.Println("ppppppppppppppppppppppppp")
		return errors.New("cannot retrieve category")
	}
	return nil
}

func (s CategoryService) ValidateRemoval(ctx context.Context, categoryId, userId int) error {
	conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d AND daps_categories.owner_id = %d", userId, categoryId, userId)
	catId, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	if err != nil {
		return err
	}
	if catId == 0 {
		return errors.New("cannot remove category")
	}
	return nil
}
