package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/services/email"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/core/services/user"
	customErrors "github.com/aghex70/daps/internal/errors"
	"github.com/aghex70/daps/internal/ports/domain"
	utils "github.com/aghex70/daps/utils/email"
	"gorm.io/gorm"
	"log"
	"time"
)

type RemindTodosUseCase struct {
	TodoService  todo.Service
	UserService  user.Service
	EmailService email.Service
	logger       *log.Logger
}

func (uc *RemindTodosUseCase) Execute(ctx context.Context) error {
	fmt.Println("Reminding users...")
	users, err := uc.UserService.List(ctx, nil, nil)
	if err != nil {
		return err
	}

	for _, user := range users {
		s, err := uc.TodoService.GetSummary(ctx, user.ID)
		if err != nil {
			return err
		}
		email, err := utils.GenerateRemindTodosHTMLContent(s)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, customErrors.ReminderAlreadySent) {
				return nil
			}
			return err
		}
		ne := domain.Email{
			From:      "daps",
			To:        user.Email,
			Recipient: user.Name,
			Subject:   fmt.Sprintf("📣 DAPS - Tareas pendientes (%s) 📣", time.Now().Format("02/01/2006")),
			Body:      email.Body,
			UserID:    user.ID,
			Source:    "daps",
		}

		sent, err := uc.EmailService.Send(ctx, ne)
		if !sent && err != nil {
			return err
		}

		return nil
	}
}
