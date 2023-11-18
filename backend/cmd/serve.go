package cmd

import (
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	emailService "github.com/aghex70/daps/internal/core/services/email"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	userService "github.com/aghex70/daps/internal/core/services/user"
	categoryUsecases "github.com/aghex70/daps/internal/core/usecases/category"
	userUsecases "github.com/aghex70/daps/internal/core/usecases/user"
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
	categoryHandler "github.com/aghex70/daps/internal/ports/handlers/category"
	userHandler "github.com/aghex70/daps/internal/ports/handlers/user"
	"log"

	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/persistence/database"
	"github.com/aghex70/daps/server"
	"github.com/spf13/cobra"
)

func ServeCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.Logger{}
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				log.Fatal("error starting database", err.Error())
			}

			// Repositories
			userr := repository.NewGormUserRepository(gdb)
			emailr := repository.NewGormEmailRepository(gdb)
			catr := repository.NewGormCategoryRepository(gdb)
			todor := repository.NewGormTodoRepository(gdb)

			//Services
			us := userService.NewUserService(userr, &logger)
			es := emailService.NewEmailService(emailr, &logger)
			cs := categoryService.NewCategoryService(catr, &logger)
			ts := todoService.NewTodoService(todor, &logger)

			// User usecases
			auuc := userUsecases.NewActivateUserUseCase(us, &logger)
			duuc := userUsecases.NewDeleteUserUseCase(us, &logger)
			guuc := userUsecases.NewGetUserUseCase(us, &logger)
			liuuc := userUsecases.NewListUsersUseCase(us, &logger)
			louuc := userUsecases.NewLoginUserUseCase(us, &logger)
			puuc := userUsecases.NewProvisionDemoUserUseCase(us, cs, ts, &logger)
			refuuc := userUsecases.NewRefreshTokenUseCase(us, &logger)
			reguuc := userUsecases.NewRegisterUserUseCase(us, cs, es, &logger)
			resuuc := userUsecases.NewResetPasswordUseCase(us, &logger)
			sruuc := userUsecases.NewSendResetLinkUseCase(us, es, &logger)
			uuuuc := userUsecases.NewUpdateUserUseCase(us, &logger)

			// Category usecases
			cauuc := categoryUsecases.NewCreateCategoryUseCase(cs, &logger)
			cduuc := categoryUsecases.NewDeleteCategoryUseCase(cs, &logger)
			gcuuc := categoryUsecases.NewGetCategoryUseCase(cs, &logger)
			lcuuc := categoryUsecases.NewListCategoriesUseCase(cs, &logger)
			scauuc := categoryUsecases.NewShareCategoryUseCase(cs, &logger)
			usauuc := categoryUsecases.NewUnshareCategoryUseCase(cs, &logger)
			ucauuc := categoryUsecases.NewUpdateCategoryUseCase(cs, &logger)

			//Handlers
			uh := userHandler.NewUserHandler(auuc, duuc, guuc, liuuc, louuc, puuc, refuuc, reguuc, resuuc, sruuc, uuuuc, &logger)
			ch := categoryHandler.NewCategoryHandler(cauuc, cduuc, gcuuc, lcuuc, scauuc, usauuc, ucauuc, &logger)
			//th := todoHandler.NewTodoHandler(ts)
			//eh := emailHandler.NewEmailHandler(es)
			//
			//rh := root.NewRootHandler(cs, ts, us)

			s := server.NewRestServer(cfg.Server.Rest, *ch, nil, *uh, nil, nil, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
