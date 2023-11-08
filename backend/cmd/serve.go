package cmd

import (
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
	categoryHandler "github.com/aghex70/daps/internal/ports/handlers/category"
	emailHandler "github.com/aghex70/daps/internal/ports/handlers/email"
	"github.com/aghex70/daps/internal/ports/handlers/root"
	todoHandler "github.com/aghex70/daps/internal/ports/handlers/todo"
	userHandler "github.com/aghex70/daps/internal/ports/handlers/user"
	"log"

	"github.com/aghex70/daps/config"
	categoryService "github.com/aghex70/daps/internal/core/usecases/category"
	emailService "github.com/aghex70/daps/internal/core/usecases/email"
	todoService "github.com/aghex70/daps/internal/core/usecases/todo"
	userService "github.com/aghex70/daps/internal/core/usecases/user"
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

			catr := repository.NewGormCategoryRepository(gdb)
			emailr := repository.NewGormEmailRepository(gdb)
			todor := repository.NewGormTodoRepository(gdb)
			userr := repository.NewGormUserRepository(gdb)

			us := userService.NewUserService(userr, &logger)
			uh := userHandler.NewUserHandler(us)

			cs := categoryService.NewCategoryService(catr, &logger)
			ch := categoryHandler.NewCategoryHandler(cs)

			ts := todoService.NewTodoService(todor, &logger)
			th := todoHandler.NewTodoHandler(ts)

			es := emailService.NewEmailService(emailr, &logger)
			eh := emailHandler.NewEmailHandler(es)

			rh := root.NewRootHandler(cs, ts, us)

			s := server.NewRestServer(cfg.Server.Rest, ch, th, uh, rh, eh, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
