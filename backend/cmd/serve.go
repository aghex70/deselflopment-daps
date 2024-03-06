package cmd

import (
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	emailService "github.com/aghex70/daps/internal/core/services/email"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	userService "github.com/aghex70/daps/internal/core/services/user"
	categoryUsecases "github.com/aghex70/daps/internal/core/usecases/category"
	todoUsecases "github.com/aghex70/daps/internal/core/usecases/todo"
	userUsecases "github.com/aghex70/daps/internal/core/usecases/user"
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
	categoryHandler "github.com/aghex70/daps/internal/ports/handlers/category"
	todoHandler "github.com/aghex70/daps/internal/ports/handlers/todo"
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
			epuc := userUsecases.NewEditProfileUseCase(us, &logger)
			guuc := userUsecases.NewGetUserUseCase(us, &logger)
			liuuc := userUsecases.NewListUsersUseCase(us, &logger)
			louuc := userUsecases.NewLoginUserUseCase(us, &logger)
			puuc := userUsecases.NewProvisionDemoUserUseCase(us, cs, ts, &logger)
			refuuc := userUsecases.NewRefreshTokenUseCase(us, &logger)
			reguuc := userUsecases.NewRegisterUserUseCase(us, cs, es, &logger)
			resuuc := userUsecases.NewResetPasswordUseCase(us, &logger)
			sruuc := userUsecases.NewSendResetLinkUseCase(us, es, &logger)

			// Category usecases
			cauuc := categoryUsecases.NewCreateCategoryUseCase(cs, us, &logger)
			cduuc := categoryUsecases.NewDeleteCategoryUseCase(cs, us, &logger)
			gcuuc := categoryUsecases.NewGetCategoryUseCase(cs, us, &logger)
			lcuuc := categoryUsecases.NewListCategoriesUseCase(cs, us, &logger)
			lcusuc := categoryUsecases.NewListCategoryUsersUseCase(cs, us, &logger)
			scauuc := categoryUsecases.NewShareCategoryUseCase(cs, us, &logger)
			usauuc := categoryUsecases.NewUnshareCategoryUseCase(cs, us, &logger)
			usuuc := categoryUsecases.NewUnsubscribeCategoryUseCase(cs, us, &logger)
			ucauuc := categoryUsecases.NewUpdateCategoryUseCase(cs, us, &logger)

			// Summary usecases
			gsuuc := categoryUsecases.NewGetSummaryUseCase(cs, us, &logger)

			// Todo usecases
			atuuc := todoUsecases.NewActivateTodoUseCase(ts, us, &logger)
			cotuuc := todoUsecases.NewCompleteTodoUseCase(ts, us, &logger)
			ctuuc := todoUsecases.NewCreateTodoUseCase(ts, us, &logger)
			dtuuc := todoUsecases.NewDeleteTodoUseCase(ts, us, &logger)
			gtuuc := todoUsecases.NewGetTodoUseCase(ts, us, &logger)
			ituuc := todoUsecases.NewImportTodosUseCase(ts, us, &logger)
			ltuuc := todoUsecases.NewListTodosUseCase(ts, us, &logger)
			rtuuc := todoUsecases.NewRestartTodoUseCase(ts, us, &logger)
			stuuc := todoUsecases.NewStartTodoUseCase(ts, us, &logger)
			utuuc := todoUsecases.NewUpdateTodoUseCase(ts, us, &logger)

			//Handlers
			uh := userHandler.NewUserHandler(auuc, duuc, epuc, guuc, liuuc, louuc, puuc, refuuc, reguuc, resuuc, sruuc, &logger)
			ch := categoryHandler.NewCategoryHandler(cauuc, cduuc, gcuuc, gsuuc, lcuuc, lcusuc, scauuc, usauuc, usuuc, ucauuc, &logger)
			th := todoHandler.NewTodoHandler(atuuc, cotuuc, ctuuc, dtuuc, gtuuc, ituuc, ltuuc, rtuuc, stuuc, utuuc, &logger)

			s := server.NewRestServer(cfg.Server.Rest, *ch, *th, *uh, &logger)
			if err = s.StartServer(); err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
