package cmd

import (
	"github.com/aghex70/daps/config"
	todoService "github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/email"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/persistence/database"
	"github.com/aghex70/daps/queues"
	"github.com/spf13/cobra"
	"log"
)

func WorkerServerCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker",
		Short: "Run worker server",
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.Logger{}
			gdb, err := database.NewGormDB(*cfg.Database)
			if err != nil {
				log.Fatal("error starting database", err.Error())
			}

			ur, _ := user.NewUserGormRepository(gdb)
			rr, _ := relationship.NewRelationshipGormRepository(gdb)
			tr, _ := todo.NewTodoGormRepository(gdb)
			er, _ := email.NewEmailGormRepository(gdb)

			tds := todoService.NewtodoService(tr, rr, er, ur, &logger)

			s := queues.NewWorkerServer(cfg.Worker, tds, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting worker server", err.Error())
			}
		},
	}
	return cmd
}
