package cmd

import (
	"github.com/aghex70/daps/config"
	categoryService "github.com/aghex70/daps/internal/core/services/category"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	userService "github.com/aghex70/daps/internal/core/services/user"
	categoryHandler "github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/root"
	todoHandler "github.com/aghex70/daps/internal/handlers/todo"
	userHandler "github.com/aghex70/daps/internal/handlers/user"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/persistence/database"
	"github.com/aghex70/daps/server"
	"github.com/spf13/cobra"
	"log"
)

func ServeCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			//log.Println("Starting app...")
			//
			//log.Println("Loading configuration...")
			//c, err := config.NewConfig()
			//log.Println("Configuration load successfully")

			// Intialize database
			//log.Println("Starting application database")
			//_, err := database.NewSqlDB(*cfg.Database)
			//if err != nil {
			//	log.Fatalf("error starting application database %+v", err.Error())
			//}

			logger2 := log.Logger{}
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				log.Fatal("error starting database", err.Error())
			}

			ur, _ := user.NewUserGormRepository(gdb)
			cr, _ := category.NewCategoryGormRepository(gdb)
			rr, _ := relationship.NewRelationshipGormRepository(gdb)
			tdr, _ := todo.NewTodoGormRepository(gdb)

			us := userService.NewUserService(ur, cr, &logger2)
			uh := userHandler.NewUserHandler(us, &logger2)

			cs := categoryService.NewCategoryService(cr, rr, &logger2)
			ch := categoryHandler.NewCategoryHandler(cs, &logger2)

			tds := todoService.NewtodoService(tdr, rr, &logger2)
			tdh := todoHandler.NewTodoHandler(tds, &logger2)

			rh := root.NewRootHandler(cs, tds, &logger2)

			s := server.NewRestServer(cfg.Server.Rest, ch, tdh, uh, rh, &logger2)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting server", err.Error())
			}
		},
	}
	return cmd
}
