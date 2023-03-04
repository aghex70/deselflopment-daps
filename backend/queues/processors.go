package queues

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/ports"
	customErrors "github.com/aghex70/daps/internal/errors"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type ReminderTodosProcessor struct {
	todoService ports.TodoServicer
}

func (processor *ReminderTodosProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p TodosReminderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := processor.todoService.Remind(nil)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, customErrors.ReminderAlreadySent) {
			// In order to avoid unnecessary retries, we return nil here.
			return nil
		}
		return err
	}
	return nil
}

func NewReminderTodosProcessor(ts ports.TodoServicer) *ReminderTodosProcessor {
	return &ReminderTodosProcessor{
		todoService: ts,
	}
}
