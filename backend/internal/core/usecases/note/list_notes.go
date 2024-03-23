package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/user"

	//"github.com/aghex70/daps/server"
	"log"
)

type ListNotesUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ListNotesUseCase) Execute(ctx context.Context, filters *map[string]interface{}, userID uint) ([]domain.Note, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Note{}, err
	}

	if !u.Active {
		return []domain.Note{}, pkg.InactiveUserError
	}

	// Set the user ID into the filters map (retrieve only own notes)
	if filters == nil {
		filters = &map[string]interface{}{}
		(*filters)["owner_id"] = userID
	} else {
		(*filters)["owner_id"] = userID
	}

	notes, err := uc.NoteService.List(ctx, filters)
	if err != nil {
		return []domain.Note{}, err
	}
	return notes, nil
}

func NewListNotesUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *ListNotesUseCase {
	return &ListNotesUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
