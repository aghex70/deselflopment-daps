package cmd

import (
	"github.com/aghex70/daps/config"

	"github.com/spf13/cobra"
)

func WorkerServerCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worker",
		Short: "Run worker server",
		Run: func(cmd *cobra.Command, args []string) {
			//logger := log.Logger{}
			//gdb, err := database.NewGormDB(*cfg.Database)
			//if err != nil {
			//	log.Fatal("error starting database", err.Error())
			//}

			//r := repository.NewGormTodoRepository(gdb)
			//
			////tds := todoService.NewTodoService(r, &logger)
			//
			//s := queues.NewWorkerServer(cfg.Broker, tds, &logger)
			//err = s.StartServer()
			//if err != nil {
			//	log.Fatal("error starting worker server", err.Error())
			//}
		},
	}
	return cmd
}
