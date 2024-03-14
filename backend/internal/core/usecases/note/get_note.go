package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/note"

	"log"
)

type GetNoteUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *GetNoteUseCase) Execute(ctx context.Context, r requests.GetNoteRequest, userID uint) (domain.Note, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Note{}, err
	}

	if !u.Active {
		return domain.Note{}, pkg.InactiveUserError
	}

	t, err := uc.NoteService.Get(ctx, r.NoteID)
	if err != nil {
		return domain.Note{}, err
	}
	if owner := utils.IsNoteOwner(t.OwnerID, u.ID); !owner {
		return domain.Note{}, pkg.UnauthorizedError
	}

	return t, nil
}

func NewGetNoteUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *GetNoteUseCase {
	return &GetNoteUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
