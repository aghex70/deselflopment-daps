package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"log"
	"mime/multipart"
)

type ImportTodosUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *ImportTodosUseCase) Execute(ctx context.Context, f multipart.File) error {
	err := uc.TodoService.Import(ctx, f)
	if err != nil {
		return err
	}
	return nil
}

func NewImportTodosUseCase(s todo.Service, logger *log.Logger) *ImportTodosUseCase {
	return &ImportTodosUseCase{
		TodoService: s,
		logger:      logger,
	}
}
