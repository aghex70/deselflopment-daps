package queues

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	customErrors "github.com/aghex70/daps/internal/common/errors"
	"github.com/aghex70/daps/internal/ports/services/todo"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type ReminderTodosProcessor struct {
	todoService todo.Servicer
}

func (processor *ReminderTodosProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p TodosReminderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := processor.todoService.Remind(context.TODO())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, customErrors.ReminderAlreadySent) {
			// In order to avoid unnecessary retries, we return nil here.
			return nil
		}
		return err
	}
	return nil
}

func NewReminderTodosProcessor(ts todo.Servicer) *ReminderTodosProcessor {
	return &ReminderTodosProcessor{
		todoService: ts,
	}
}
