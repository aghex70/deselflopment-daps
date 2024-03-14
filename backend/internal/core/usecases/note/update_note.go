package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/note"

	"log"
)

type UpdateNoteUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *UpdateNoteUseCase) Execute(ctx context.Context, r requests.UpdateNoteRequest, userID uint) error {
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
	owner := utils.IsNoteOwner(t.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	fields := common.StructToMap(r, "note_id")
	if err = uc.NoteService.Update(ctx, t.ID, &fields); err != nil {
		return err
	}
	return nil
}

func NewUpdateNoteUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
