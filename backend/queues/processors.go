package queues

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/hibiken/asynq"
)

type ReminderTodosProcessor struct {
	todoService ports.TodoServicer
}

func (processor *ReminderTodosProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p TodosReminderPayload
	fmt.Println("@@@@@@@@@@@@@@@@ Processing task: ", t.Type())
	fmt.Println("@@@@@@@@@@@@@@@@ Processing task: ", t.Type())
	fmt.Println("@@@@@@@@@@@@@@@@ Processing task: ", t.Type())
	fmt.Println("@@@@@@@@@@@@@@@@ Processing task: ", t.Type())
	fmt.Println("@@@@@@@@@@@@@@@@ Processing task: ", t.Type())
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	err := processor.todoService.Remind(nil)
	if err != nil {
		return err
	}
	return nil
}

func NewReminderTodosProcessor(ts ports.TodoServicer) *ReminderTodosProcessor {
	return &ReminderTodosProcessor{
		todoService: ts,
	}
}
