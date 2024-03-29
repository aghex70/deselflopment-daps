package todo

import (
	"bufio"
	"context"
	"encoding/csv"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/todo"
	"io"
	"log"
	"mime/multipart"
	"strconv"
)

type Service struct {
	logger         *log.Logger
	todoRepository todo.Repository
}

func (s Service) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	t, err := s.todoRepository.Create(ctx, t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.Todo, error) {
	t, err := s.todoRepository.Get(ctx, id)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	if err := s.todoRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Todo, error) {
	todos, err := s.todoRepository.List(ctx, filters)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s Service) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	if err := s.todoRepository.Update(ctx, id, filters); err != nil {
		return err
	}
	return nil
}

func (s Service) Start(ctx context.Context, id uint) error {
	if err := s.todoRepository.Start(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) Complete(ctx context.Context, id uint) error {
	if err := s.todoRepository.Complete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) Restart(ctx context.Context, id uint) error {
	if err := s.todoRepository.Restart(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) Activate(ctx context.Context, id uint) error {
	if err := s.todoRepository.Activate(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) Import(ctx context.Context, f multipart.File) error {
	// Create a buffer to read the file line by line
	buf := bufio.NewReader(f)

	// Parse the CSV file
	rr := csv.NewReader(buf)

	// Read and discard the first line
	_, err := rr.Read()
	if err != nil {
		return err
	}

	// Iterate over the lines of the CSV file
	for {
		// Read the next line
		record, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		name := record[0]
		link := record[1]
		categoryID, _ := strconv.Atoi(record[2])

		_, err = s.todoRepository.Create(ctx, domain.Todo{
			Name:       name,
			Link:       &link,
			CategoryID: uint(categoryID),
			Priority:   domain.Priority(3),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//func (s Service) Suggest(ctx context.Context, r *http.Request) error {
//	//userId, _ := server.RetrieveJWTClaims(r, nil)
//	err := s.todoRepository.Suggest(ctx, int(userId))
//	if err != nil {
//		return err
//	}
//	return nil
//}

func NewTodoService(tr todo.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		todoRepository: tr,
	}
}
