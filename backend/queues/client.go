package queues

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aghex70/daps/config"
	"github.com/hibiken/asynq"
)

type WorkerClient struct {
	logger *log.Logger
	cfg    config.CacheConfig
}

func (s *WorkerClient) StartClient() (*asynq.Client, error) {
	port := strconv.Itoa(s.cfg.Port)
	address := fmt.Sprintf("%s:%s", s.cfg.Host, port)
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: address})
	return client, nil
}

func NewWorkerClient(cfg *config.CacheConfig, logger *log.Logger) *WorkerClient {
	return &WorkerClient{
		cfg:    *cfg,
		logger: logger,
	}
}
