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
	err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, ids *[]uint, fields *map[string]interface{}) ([]domain.Todo, error) {
	todos, err := s.todoRepository.List(ctx, ids, fields)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s Service) Update(ctx context.Context, id uint, fields *map[string]interface{}) error {
	err := s.todoRepository.Update(ctx, id, fields)
	if err != nil {
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

//func (s Service) GetSummary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error) {
//	//userId, _ := server.RetrieveJWTClaims(r, nil)
//	summary, err := s.todoRepository.GetSummary(ctx, int(userId))
//	if err != nil {
//		return []domain.CategorySummary{}, err
//	}
//	return summary, nil
//}

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
