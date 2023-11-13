package category

import (
	"context"
)

func (s Service) ValidateCreation(ctx context.Context, name string, userID uint) error {
	//conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_categories.name = '%s'", userID, name)
	//categoryID, err := s.r.UserCategoryExists(ctx, conditions)
	//if err != nil {
	//	return err
	//}
	//if categoryID != 0 {
	//	return errors.New("already existent category with that user and name")
	//}
	return nil
}

func (s Service) ValidateModification(ctx context.Context, categoryID, userID uint) error {
	//conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d AND daps_categories.owner_id = %d", userID, categoryID, userID)
	//catID, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	//if err != nil {
	//	return err
	//}
	//if catID == 0 {
	//	return errors.New("cannot update category")
	//}
	return nil
}

func (s Service) ValidateShare(ctx context.Context, categoryID, userID uint) error {
	//conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d AND daps_categories.owner_id = %d", userID, categoryID, userID)
	//catID, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	//if err != nil {
	//	return err
	//}
	//if catID == 0 {
	//	return errors.New("cannot update category")
	//}
	return nil
}

func (s Service) ValidateUnshare(ctx context.Context, categoryID, userID uint) error {
	//conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d", userID, categoryID)
	//catID, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	//if err != nil {
	//	return err
	//}
	//if catID == 0 {
	//	return errors.New("cannot update category")
	//}
	return nil
}

func (s Service) ValidateRetrieval(ctx context.Context, categoryID, userID uint) error {
	//conditions := fmt.Sprintf("daps_category_users.user_id = %d AND daps_category_users.category_id = %d", userID, categoryID)
	//catID, err := s.categoryRepository.UserCategoryExists(ctx, conditions)
	//if err != nil {
	//	return err
	//}
	//if catID == 0 {
	//	return errors.New("cannot retrieve category")
	//}
	return nil
}
