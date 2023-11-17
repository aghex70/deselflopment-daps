package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	common "github.com/aghex70/daps/utils"
	categoryUtils "github.com/aghex70/daps/utils/category"
	userUtils "github.com/aghex70/daps/utils/user"
	"log"
)

type ProvisionDemoUserUseCase struct {
	UserService     user.Service
	CategoryService category.Service
	TodoService     todo.Service
	logger          *log.Logger
}

func (uc *ProvisionDemoUserUseCase) Execute(ctx context.Context, r requests.ProvisionDemoUserRequest, userID uint) error {
	ru, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}
	if !ru.Admin {
		return pkg.UnauthorizedError
	}
	encryptedPassword := userUtils.EncryptPassword(ctx, r.Password)

	u := domain.User{
		Name:              r.Name,
		Email:             r.Email,
		Password:          encryptedPassword,
		Active:            false,
		ActivationCode:    common.GenerateUUID(),
		ResetPasswordCode: common.GenerateUUID(),
	}

	nu, err := uc.UserService.Create(ctx, u)
	if err != nil {
		return err
	}

	demoCategory := domain.Category{
		OwnerID: nu.ID,
		//Description: "Home tasks",
		Custom: true,
		Name:   "Home",
		Users:  &[]domain.User{u},
	}

	c, err := uc.CategoryService.Create(ctx, demoCategory)
	if err != nil {
		return err
	}

	anotherDemoCategory := domain.Category{
		OwnerID: nu.ID,
		//Description: "Work stuff",
		Custom: true,
		Name:   "Work",
		Users:  &[]domain.User{u},
	}

	ac, err := uc.CategoryService.Create(ctx, anotherDemoCategory)
	if err != nil {
		return err
	}

	yetAnotherDemoCategory := domain.Category{
		OwnerID: nu.ID,
		//Description: "Purchase list",
		Custom: true,
		Name:   "Purchases",
		Users:  &[]domain.User{u},
	}

	yac, err := uc.CategoryService.Create(ctx, yetAnotherDemoCategory)
	if err != nil {
		return err
	}

	todos := categoryUtils.GenerateDemoTodos(c.ID, ac.ID, yac.ID, nu.ID, "es")

	for _, t := range todos {
		_, err = uc.TodoService.Create(ctx, t)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewProvisionDemoUserUseCase(userService user.Service, categoryService category.Service, todoService todo.Service, logger *log.Logger) *ProvisionDemoUserUseCase {
	return &ProvisionDemoUserUseCase{
		UserService:     userService,
		CategoryService: categoryService,
		TodoService:     todoService,
		logger:          logger,
	}
}
