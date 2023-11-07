package todo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	customErrors "github.com/aghex70/daps/internal/errors"
	repository "github.com/aghex70/daps/internal/repositories/gorm"
	"github.com/aghex70/daps/pkg"
	"github.com/aghex70/daps/server"
	"gorm.io/gorm"
)

type Service struct {
	logger     *log.Logger
	repository *repository.GormRepository
}

func (s Service) Create(ctx context.Context, r *http.Request, req ports.CreateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	if req.Category != 1 {
		err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
		if err != nil {
			return err
		}
	}
	_, preexistent := s.CheckExistentTodo(ctx, req.Name, req.Category)
	if preexistent {
		return errors.New("already existent todo with that info")
	}

	ntd := domain.Todo{
		Category:    req.Category,
		Description: req.Description,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
		Recurring:   req.Recurring,
		Recurrency:  req.Recurrency,
		Suggestable: true,
	}

	err, _ := s.repository.CreateTodo(ctx, ntd)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Update(ctx context.Context, r *http.Request, req ports.UpdateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}

	ntd := domain.Todo{
		Id:          int(req.TodoId),
		Category:    req.Category,
		Description: req.Description,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
		Recurring:   req.Recurring,
		Recurrency:  req.Recurrency,
		Suggestable: req.Suggestable,
	}

	err = s.repository.UpdateTodo(ctx, ntd)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Complete(ctx context.Context, r *http.Request, req ports.CompleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.repository.Complete(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Activate(ctx context.Context, r *http.Request, req ports.ActivateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.repository.Activate(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Start(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.repository.Start(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Restart(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.repository.Restart(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request, req ports.GetTodoRequest) (domain.TodoInfo, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return domain.TodoInfo{}, err
	}
	td, err := s.repository.GetTodo(ctx, int(req.TodoId))
	if err != nil {
		return domain.TodoInfo{}, err
	}
	return td, nil
}

func (s Service) List(ctx context.Context, r *http.Request, req ports.ListTodosRequest) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.repository.GetTodos(ctx, req.Category)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s Service) ListRecurring(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.repository.GetTodos(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s Service) ListCompleted(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.repository.GetTodos(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s Service) ListSuggested(ctx context.Context, r *http.Request) ([]domain.TodoInfo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	todos, err := s.repository.GetTodos(ctx, int(userId))
	if err != nil {
		return []domain.TodoInfo{}, err
	}
	return todos, nil
}

func (s Service) Suggest(ctx context.Context, r *http.Request) error {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	err := s.repository.Suggest(ctx, int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req ports.DeleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.repository.DeleteTodo(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Summary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	summary, err := s.repository.GetSummary(ctx, int(userId))
	if err != nil {
		return []domain.CategorySummary{}, err
	}
	return summary, nil
}

func (s Service) Remind(ctx context.Context) error {
	fmt.Println("Reminding users...")
	users, err := s.repository.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, u := range users {
		rs, err := s.repository.GetRemindSummary(ctx, u.Id)
		fmt.Printf("Remind summary for user %d: %+v", u.Id, rs)
		if err != nil {
			return err
		}
		e, err := pkg.GenerateRemindTodosHTMLContent(u, rs)
		if err != nil {
			if err == gorm.ErrRecordNotFound || err == customErrors.ReminderAlreadySent {
				return nil
			}
			return err
		}

		ne := domain.Email{
			From:      pkg.FromEmail,
			To:        u.Email,
			Recipient: u.Name,
			Subject:   fmt.Sprintf("📣 DAPS - Tareas pendientes (%s) 📣", time.Now().Format("02/01/2006")),
			Body:      e.Body,
			User:      u.Id,
			Source:    "daps",
		}

		err = pkg.SendEmail(ne)
		if err != nil {
			fmt.Printf("Error sending email: %+v", err)
			ne.Error = err.Error()
			ne.Sent = false
			_, errz := s.repository.CreateEmail(ctx, ne)
			if errz != nil {
				fmt.Printf("Error saving email: %+v", errz)
				return errz
			}
			return err
		}

		ne.Sent = true
		_, err = s.repository.CreateEmail(ctx, ne)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewTodoService(r *repository.GormRepository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
