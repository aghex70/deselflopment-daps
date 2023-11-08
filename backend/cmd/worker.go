package cmd

import (
	repository "github.com/aghex70/daps/internal/infrastructure/persistence/repositories/gorm"
	"log"

	"github.com/aghex70/daps/config"
	todoService "github.com/aghex70/daps/internal/core/usecases/todo"
	"github.com/aghex70/daps/persistence/database"
	"github.com/aghex70/daps/queues"
	"github.com/spf13/cobra"
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

			r := repository.NewGormTodoRepository(gdb)

			tds := todoService.NewTodoService(r, &logger)

			s := queues.NewWorkerServer(cfg.Broker, tds, &logger)
			err = s.StartServer()
			if err != nil {
				log.Fatal("error starting worker server", err.Error())
			}
		},
	}
	return cmd
}
