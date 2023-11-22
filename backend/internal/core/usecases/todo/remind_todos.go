package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/services/email"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type RemindTodosUseCase struct {
	TodoService  todo.Servicer
	UserService  user.Servicer
	EmailService email.Servicer
	logger       *log.Logger
}

func (uc *RemindTodosUseCase) Execute(ctx context.Context) error {
	//fmt.Println("Reminding users...")
	//users, err := uc.UserService.List(ctx, nil, nil)
	//if err != nil {
	//	return err
	//}
	//
	//for _, user := range users {
	//	s, err := uc.TodoService.GetSummary(ctx, user.ID)
	//	if err != nil {
	//		return err
	//	}
	//	email, err := utils.GenerateRemindTodosHTMLContent(s)
	//	if err != nil {
	//		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, customErrors.ReminderAlreadySent) {
	//			return nil
	//		}
	//		return err
	//	}
	//	ne := domain.Email{
	//		From:      "daps",
	//		To:        user.Email,
	//		Recipient: user.Name,
	//		Subject:   fmt.Sprintf("📣 DAPS - Tareas pendientes (%s) 📣", time.Now().Format("02/01/2006")),
	//		Body:      email.Body,
	//		OwnerID:    user.ID,
	//		Source:    "daps",
	//	}
	//
	//	sent, err := uc.EmailService.Send(ctx, ne)
	//	if !sent && err != nil {
	//		return err
	//	}
	//
	//	return nil
	//}
	return nil
}
