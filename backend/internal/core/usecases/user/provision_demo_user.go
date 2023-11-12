package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"net/http"
)

func (s Service) ProvisionDemoUser(ctx context.Context, r *http.Request, req requests.ProvisionDemoUserRequest) error {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//cipheredPassword := s.EncryptPassword(ctx, req.Password)
	//u := domain2.User{
	//	Name:              pkg2.DemoUserName,
	//	Email:             req.Email,
	//	Password:          cipheredPassword,
	//	Active:            true,
	//	ResetPasswordCode: pkg2.GenerateUUID(),
	//}
	//
	//nu, err := s.repository.CreateUser(ctx, u)
	//if err != nil {
	//	return err
	//}
	//
	////nuc := domain.UserConfig{
	////	UserID:      nu.ID,
	////	AutoSuggest: false,
	////	Language:    "en",
	////}
	//
	////err = s.userConfigurationRepository.Create(ctx, nuc)
	////err = s.userConfigurationRepository.Create(ctx, nuc)
	////if err != nil {
	////	return err
	////}
	//
	//demoCategory := domain2.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Home tasks",
	//	Custom: true,
	//	Name:   "Home",
	//	//Users:       []domain.User{u},
	//}
	//
	//c, err := s.repository.CreateCategory(ctx, demoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//anotherDemoCategory := domain2.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Work stuff",
	//	Custom: true,
	//	Name:   "Work",
	//	//Users:       []domain.User{u},
	//}
	//
	//ac, err := s.repository.CreateCategory(ctx, anotherDemoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//yetAnotherDemoCategory := domain2.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Purchase list",
	//	Custom: true,
	//	Name:   "Purchases",
	//	//Users:       []domain.User{u},
	//}
	//
	//yac, err := s.repository.CreateCategory(ctx, yetAnotherDemoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//todos := pkg2.GenerateDemoTodos(int(c.ID), int(ac.ID), int(yac.ID), req.Language)
	//
	//for _, t := range todos {
	//	_, _ = s.repository.CreateTodo(ctx, t)
	//	//if err != nil {
	//	//	return err
	//	//}
	//}

	return nil
}
