package queues

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/hibiken/asynq"
	"log"
	"strconv"
)

type WorkerClient struct {
	logger *log.Logger
	cfg    config.WorkerConfig
}

func (s *WorkerClient) StartClient() (*asynq.Client, error) {
	port := strconv.Itoa(s.cfg.BrokerConfig.Port)
	address := fmt.Sprintf("%s:%s", s.cfg.BrokerConfig.Host, port)
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: address})
	return client, nil
}

func NewWorkerClient(cfg *config.WorkerConfig, logger *log.Logger) *WorkerClient {
	return &WorkerClient{
		cfg:    *cfg,
		logger: logger,
	}
}
