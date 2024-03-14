package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type CreateNoteUseCase struct {
	NoteService  note.Servicer
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
}

func (uc *CreateNoteUseCase) Execute(ctx context.Context, userID uint, r requests.CreateNoteRequest) (domain.Note, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Note{}, err
	}

	if !u.Active {
		return domain.Note{}, pkg.InactiveUserError
	}
	to, err := uc.TopicService.Get(ctx, r.TopicID)
	if err != nil {
		return domain.Note{}, err
	}

	t := domain.Note{
		Content: r.Content,
		Topics:  []domain.Topic{to},
		Shared:  false,
		OwnerID: u.ID,
		Users:   []domain.User{u},
	}
	t, err = uc.NoteService.Create(ctx, t)
	if err != nil {
		return domain.Note{}, err
	}

	return t, nil
}

func NewCreateNoteUseCase(s note.Servicer, u user.Servicer, t topic.Servicer, logger *log.Logger) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		NoteService:  s,
		UserService:  u,
		TopicService: t,
		logger:       logger,
	}
}
