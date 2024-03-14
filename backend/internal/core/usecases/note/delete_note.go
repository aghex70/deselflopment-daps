package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/note"

	"log"
)

type DeleteNoteUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *DeleteNoteUseCase) Execute(ctx context.Context, r requests.DeleteNoteRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	t, err := uc.NoteService.Get(ctx, r.NoteID)
	if err != nil {
		return err
	}
	owner := utils.IsNoteOwner(t.OwnerID, u.ID)
	if !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.NoteService.Delete(ctx, r.NoteID); err != nil {
		return err
	}
	return nil
}

func NewDeleteNoteUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *DeleteNoteUseCase {
	return &DeleteNoteUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
