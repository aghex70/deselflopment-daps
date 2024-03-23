package note

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/note"
	"log"
)

type Service struct {
	logger         *log.Logger
	noteRepository note.Repository
}

func (s Service) Create(ctx context.Context, n domain.Note) (domain.Note, error) {
	nn, err := s.noteRepository.Create(ctx, n)
	if err != nil {
		return nn, err
	}
	return nn, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.Note, error) {
	n, err := s.noteRepository.Get(ctx, id)
	fmt.Printf("Service.Get: %+v\n", n)
	if err != nil {
		return domain.Note{}, err
	}
	return n, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	if err := s.noteRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Note, error) {
	notes, err := s.noteRepository.List(ctx, filters)
	if err != nil {
		return []domain.Note{}, err
	}
	return notes, nil
}

func (s Service) Update(ctx context.Context, id uint, fields *map[string]interface{}) error {
	if err := s.noteRepository.Update(ctx, id, fields); err != nil {
		return err
	}
	return nil
}

func (s Service) Share(ctx context.Context, id uint, u domain.User) error {
	if err := s.noteRepository.Share(ctx, id, u); err != nil {
		return err
	}
	return nil
}

func (s Service) Unshare(ctx context.Context, id uint, u domain.User) error {
	if err := s.noteRepository.Unshare(ctx, id, u); err != nil {
		return err
	}
	return nil
}

func NewNoteService(tr note.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		noteRepository: tr,
	}
}
