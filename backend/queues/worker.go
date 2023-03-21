package queues

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/hibiken/asynq"
)

type WorkerServer struct {
	logger      *log.Logger
	cfg         config.CacheConfig
	todoService ports.TodoServicer
}

func (s *WorkerServer) StartServer() error {
	port := strconv.Itoa(s.cfg.Port)
	address := fmt.Sprintf("%s:%s", s.cfg.Host, port)
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: address},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 4,
			// Optionally specify multiple queues with different priority.
			//Queues: map[string]int{
			//	"critical": 6,
			//	"default":  3,
			//	"low":      1,
			// },
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	processor := NewReminderTodosProcessor(s.todoService)
	mux.Handle(TypeTodosReminder, processor)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
	return nil
}

func NewWorkerServer(cfg *config.CacheConfig, ts ports.TodoServicer, logger *log.Logger) *WorkerServer {
	return &WorkerServer{
		cfg:         *cfg,
		logger:      logger,
		todoService: ts,
	}
}
