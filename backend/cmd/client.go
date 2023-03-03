package cmd

import (
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/queues"
	"github.com/spf13/cobra"
	"log"
)

func WorkerClientCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Run worker client",
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.Logger{}
			wc := queues.NewWorkerClient(cfg.Worker, &logger)
			c, err := wc.StartClient()
			if err != nil {
				log.Fatal("error starting worker client", err.Error())
			}
			defer c.Close()
			task, err := queues.NewTodosReminderTask()
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}
			info, err := c.Enqueue(task)
			if err != nil {
				log.Fatalf("could not enqueue task: %v", err)
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

		},
	}
	return cmd
}
