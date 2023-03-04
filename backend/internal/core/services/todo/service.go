package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	customErrors "github.com/aghex70/daps/internal/errors"
	"github.com/aghex70/daps/internal/repositories/gorm/email"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/pkg"
	"github.com/aghex70/daps/server"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type TodoService struct {
	logger                 *log.Logger
	todoRepository         *todo.TodoGormRepository
	relationshipRepository *relationship.RelationshipGormRepository
	emailRepository        *email.EmailGormRepository
	userRepository         *user.UserGormRepository
}

func (s TodoService) Create(ctx context.Context, r *http.Request, req ports.CreateTodoRequest) error {
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

	err := s.todoRepository.Create(ctx, ntd)
	if err != nil {
		return err
	}

	return nil
}

func (s TodoService) Update(ctx context.Context, r *http.Request, req ports.UpdateTodoRequest) error {
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

	err = s.todoRepository.Update(ctx, ntd)
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Complete(ctx context.Context, r *http.Request, req ports.CompleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Complete(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Activate(ctx context.Context, r *http.Request, req ports.ActivateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Activate(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Start(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Start(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Restart(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Restart(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Get(ctx context.Context, r *http.Request, req ports.GetTodoRequest) (domain.TodoInfo, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return domain.TodoInfo{}, err
	}
	td, err := s.todoRepository.GetById(ctx, int(req.TodoId))
	if err != nil {
		return domain.TodoInfo{}, err
	}
	return td, nil
}

func (s TodoService) List(ctx context.Context, r *http.Request, req ports.ListTodosRequest) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.List(ctx, req.Category)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) ListRecurring(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.ListRecurring(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) ListCompleted(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.ListCompleted(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) ListSuggested(ctx context.Context, r *http.Request) ([]domain.TodoInfo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	todos, err := s.todoRepository.ListSuggested(ctx, int(userId))
	if err != nil {
		return []domain.TodoInfo{}, err
	}
	return todos, nil
}

func (s TodoService) Suggest(ctx context.Context, r *http.Request) error {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	err := s.todoRepository.Suggest(ctx, int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Delete(ctx context.Context, r *http.Request, req ports.DeleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Delete(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Summary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	summary, err := s.todoRepository.GetSummary(ctx, int(userId))
	if err != nil {
		return []domain.CategorySummary{}, err
	}
	return summary, nil
}

func (s TodoService) Remind(ctx context.Context) error {
	fmt.Println("Reminding users...")
	users, err := s.userRepository.List(ctx)
	fmt.Printf("Users: %+v", users)
	if err != nil {
		fmt.Printf("Error retrieving users: %+v", err)
		return err
	}

	for _, u := range users {
		rs, err := s.todoRepository.GetRemindSummary(ctx, u.Id)
		fmt.Printf("Remind summary for user %d: %+v", u.Id, rs)
		if err != nil {
			return err
		}
		e, err := pkg.GenerateRemindTodosHTMLContent(u, rs)
		if err != nil {
			fmt.Printf("Error generating email content: %+v", err)
			fmt.Printf("Error generating email content: %+v", err)
			fmt.Printf("Error generating email content: %+v", err)
			fmt.Printf("Error generating email content: %+v", err)
			fmt.Printf("Error generating email content: %+v", err)
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
		}

		err = pkg.SendEmail(ne)
		if err != nil {
			fmt.Printf("Error sending email: %+v", err)
			ne.Error = err.Error()
			ne.Sent = false
			_, errz := s.emailRepository.Create(ctx, ne)
			if errz != nil {
				fmt.Printf("Error saving email: %+v", errz)
				return errz
			}
			return err
		}

		ne.Sent = true
		_, err = s.emailRepository.Create(ctx, ne)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewtodoService(tr *todo.TodoGormRepository, rr *relationship.RelationshipGormRepository, er *email.EmailGormRepository, ur *user.UserGormRepository, logger *log.Logger) TodoService {
	return TodoService{
		logger:                 logger,
		todoRepository:         tr,
		relationshipRepository: rr,
		emailRepository:        er,
		userRepository:         ur,
	}
}
