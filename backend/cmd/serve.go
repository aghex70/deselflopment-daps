package cmd

import (
	"log"

	"github.com/aghex70/daps/config"
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	userService "github.com/aghex70/daps/internal/core/services/user"
	categoryHandler "github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/root"
	todoHandler "github.com/aghex70/daps/internal/handlers/todo"
	userHandler "github.com/aghex70/daps/internal/handlers/user"
	repository "github.com/aghex70/daps/internal/repositories/gorm"
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

			r, _ := repository.NewGormRepository(gdb)

			us := userService.NewUserService(r, &logger)
			uh := userHandler.NewUserHandler(us)

			cs := categoryService.NewCategoryService(r, &logger)
			ch := categoryHandler.NewCategoryHandler(cs)

			ts := todoService.NewTodoService(r, &logger)
			th := todoHandler.NewTodoHandler(ts)

			es := emailService.NewEmailService(r, &logger)
			eh := emailHandler.NewEmailHandler(es)

			rh := root.NewRootHandler(cs, ts, us)

			s := server.NewRestServer(cfg.Server.Rest, ch, th, uh, rh, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
