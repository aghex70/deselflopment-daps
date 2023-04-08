package cmd

import (
	"log"

	"github.com/aghex70/daps/config"
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	userService "github.com/aghex70/daps/internal/core/services/user"
	userConfigService "github.com/aghex70/daps/internal/core/services/userconfig"
	categoryHandler "github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/root"
	todoHandler "github.com/aghex70/daps/internal/handlers/todo"
	userHandler "github.com/aghex70/daps/internal/handlers/user"
	userConfigHandler "github.com/aghex70/daps/internal/handlers/userconfig"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/email"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/internal/repositories/gorm/userconfig"
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

			ur, _ := user.NewUserGormRepository(gdb)
			cr, _ := category.NewGormRepository(gdb)
			rr, _ := relationship.NewRelationshipGormRepository(gdb)
			tr, _ := todo.NewTodoGormRepository(gdb)
			ucr, _ := userconfig.NewUserConfigGormRepository(gdb)
			er, _ := email.NewEmailGormRepository(gdb)

			us := userService.NewUserService(ur, cr, ucr, tr, er, &logger)
			uh := userHandler.NewUserHandler(us)

			cs := categoryService.NewCategoryService(cr, rr, er, &logger)
			ch := categoryHandler.NewCategoryHandler(cs)

			tds := todoService.NewtodoService(tr, rr, er, ur, &logger)
			tdh := todoHandler.NewTodoHandler(tds)

			ucs := userConfigService.NewUserConfigService(ucr, &logger)
			uch := userConfigHandler.NewUserConfigHandler(ucs)

			rh := root.NewRootHandler(cs, tds, us)

			s := server.NewRestServer(cfg.Server.Rest, ch, tdh, uh, rh, uch, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
