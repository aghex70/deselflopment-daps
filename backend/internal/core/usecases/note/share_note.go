package note

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/note"
	"log"
)

type ShareNoteUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ShareNoteUseCase) Execute(ctx context.Context, r requests.ShareNoteRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	nu, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		fmt.Printf("Error getting user by email: %v\n", err)
		return err
	}

	c, err := uc.NoteService.Get(ctx, r.NoteID)
	if err != nil {
		return err
	}
	if owner := utils.IsNoteOwner(c.OwnerID, userID); !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.NoteService.Share(ctx, c.ID, nu); err != nil {
		return err
	}
	return nil
}

func NewShareNoteUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *ShareNoteUseCase {
	return &ShareNoteUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
