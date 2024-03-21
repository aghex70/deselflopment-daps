package cmd

import (
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	emailService "github.com/aghex70/daps/internal/core/services/email"
	noteService "github.com/aghex70/daps/internal/core/services/note"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	topicService "github.com/aghex70/daps/internal/core/services/topic"
	userService "github.com/aghex70/daps/internal/core/services/user"
	categoryUsecases "github.com/aghex70/daps/internal/core/usecases/category"
	noteUsecases "github.com/aghex70/daps/internal/core/usecases/note"
	todoUsecases "github.com/aghex70/daps/internal/core/usecases/todo"
	topicUsecases "github.com/aghex70/daps/internal/core/usecases/topic"
	userUsecases "github.com/aghex70/daps/internal/core/usecases/user"
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
	categoryHandler "github.com/aghex70/daps/internal/ports/handlers/category"
	noteHandler "github.com/aghex70/daps/internal/ports/handlers/note"
	todoHandler "github.com/aghex70/daps/internal/ports/handlers/todo"
	topicHandler "github.com/aghex70/daps/internal/ports/handlers/topic"
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
			topicr := repository.NewGormTopicRepository(gdb)
			noter := repository.NewGormNoteRepository(gdb)

			//Services
			us := userService.NewUserService(userr, &logger)
			es := emailService.NewEmailService(emailr, &logger)
			cs := categoryService.NewCategoryService(catr, &logger)
			ts := todoService.NewTodoService(todor, &logger)
			tos := topicService.NewTopicService(topicr, &logger)
			ns := noteService.NewNoteService(noter, &logger)

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

			// Checklist usecases
			gcluuc := todoUsecases.NewGetChecklistUseCase(ts, us, &logger)

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

			// Topic usecases
			ctouuc := topicUsecases.NewCreateTopicUseCase(tos, us, &logger)
			dtouuc := topicUsecases.NewDeleteTopicUseCase(tos, us, &logger)
			gtouuc := topicUsecases.NewGetTopicUseCase(tos, us, &logger)
			ltouuc := topicUsecases.NewListTopicsUseCase(tos, us, &logger)
			utouuc := topicUsecases.NewUpdateTopicUseCase(tos, us, &logger)

			// Note usecases
			cnuuc := noteUsecases.NewCreateNoteUseCase(ns, us, tos, &logger)
			dnuuc := noteUsecases.NewDeleteNoteUseCase(ns, us, &logger)
			gnuuc := noteUsecases.NewGetNoteUseCase(ns, us, &logger)
			lnuuc := noteUsecases.NewListNotesUseCase(ns, us, &logger)
			snuuc := noteUsecases.NewShareNoteUseCase(ns, us, &logger)
			usnuuc := noteUsecases.NewUnshareNoteUseCase(ns, us, &logger)
			unuuc := noteUsecases.NewUpdateNoteUseCase(ns, us, &logger)

			//Handlers
			uh := userHandler.NewUserHandler(auuc, duuc, epuc, guuc, liuuc, louuc, puuc, refuuc, reguuc, resuuc, sruuc, &logger)
			ch := categoryHandler.NewCategoryHandler(cauuc, cduuc, gcuuc, gsuuc, lcuuc, lcusuc, scauuc, usauuc, usuuc, ucauuc, &logger)
			th := todoHandler.NewTodoHandler(atuuc, cotuuc, ctuuc, dtuuc, gcluuc, gtuuc, ituuc, ltuuc, rtuuc, stuuc, utuuc, &logger)
			toh := topicHandler.NewTopicHandler(ctouuc, dtouuc, gtouuc, ltouuc, utouuc, &logger)
			nh := noteHandler.NewNoteHandler(cnuuc, dnuuc, gnuuc, lnuuc, snuuc, usnuuc, unuuc, &logger)

			s := server.NewRestServer(cfg.Server.Rest, *ch, *nh, *th, *toh, *uh, &logger)
			if err = s.StartServer(); err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
