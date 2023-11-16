package cmd

import (
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	emailService "github.com/aghex70/daps/internal/core/services/email"
	userService "github.com/aghex70/daps/internal/core/services/user"
	userUsecases "github.com/aghex70/daps/internal/core/usecases/user"
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
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
			//todor := repository.NewGormTodoRepository(gdb)

			//Services
			us := userService.NewUserService(userr, &logger)
			es := emailService.NewEmailService(emailr, &logger)
			cs := categoryService.NewCategoryService(catr, &logger)
			//ts := todoService.NewTodoService(todor, &logger)

			//UseCases
			auuc := userUsecases.NewActivateUserUseCase(us, &logger)
			//cauuc := userUsecases.NewCheckAdminUseCase(us, &logger)
			//duuc := userUsecases.NewDeleteUserUseCase(us, &logger)
			//guuc := userUsecases.NewGetUserUseCase(us, &logger)
			//liuuc := userUsecases.NewListUsersUseCase(us, &logger)
			louuc := userUsecases.NewLoginUserUseCase(us, &logger)
			//puuc := userUsecases.NewProvisionDemoUserUseCase(us, cs, ts, &logger)
			refuuc := userUsecases.NewRefreshTokenUseCase(us, &logger)
			reguuc := userUsecases.NewRegisterUserUseCase(us, cs, es, &logger)
			resuuc := userUsecases.NewResetPasswordUseCase(us, &logger)
			sruuc := userUsecases.NewSendResetLinkUseCase(us, es, &logger)

			//Handlers
			uh := userHandler.NewUserHandler(auuc, louuc, refuuc, reguuc, resuuc, sruuc, &logger)
			//ch := categoryHandler.NewCategoryHandler(cs)
			//th := todoHandler.NewTodoHandler(ts)
			//eh := emailHandler.NewEmailHandler(es)
			//
			//rh := root.NewRootHandler(cs, ts, us)

			s := server.NewRestServer(cfg.Server.Rest, nil, nil, *uh, nil, nil, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
