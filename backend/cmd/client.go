package cmd

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/queues"
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"time"
)

func WorkerClientCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Run worker client",
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.Logger{}
			wc := queues.NewWorkerClient(cfg.Cache, &logger)
			c, err := wc.StartClient()
			if err != nil {
				log.Fatal("error starting worker client", err.Error())
			}
			defer c.Close()

			location, err := time.LoadLocation("Europe/Madrid")
			if err != nil {
				panic(err)
			}

			port := strconv.Itoa(cfg.Cache.Port)
			address := fmt.Sprintf("%s:%s", cfg.Cache.Host, port)
			scheduler := asynq.NewScheduler(
				asynq.RedisClientOpt{Addr: address},
				&asynq.SchedulerOpts{
					Location: location,
				},
			)

			task, err := queues.NewTodosReminderTask()
			if err != nil {
				log.Fatalf("could not create task: %v", err)
			}

			entryID, err := scheduler.Register("30 7 * * *", task)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("registered an entry: %q\n", entryID)

			if err := scheduler.Run(); err != nil {
				log.Fatal(err)
			}

		},
	}
	return cmd
}
