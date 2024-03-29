package queues

import (
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
)

func NewTodosReminderTask() (*asynq.Task, error) {
	payload, err := json.Marshal(TodosReminderPayload{})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeTodosReminder, payload, asynq.MaxRetry(5), asynq.Timeout(5*time.Minute)), nil
}
