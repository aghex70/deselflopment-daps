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

type UnshareNoteUseCase struct {
	NoteService note.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *UnshareNoteUseCase) Execute(ctx context.Context, r requests.UnshareNoteRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	du, err := uc.UserService.GetByEmail(ctx, r.Email)
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

	if err = uc.NoteService.Unshare(ctx, c.ID, du); err != nil {
		return err
	}
	return nil
}

func NewUnshareNoteUseCase(s note.Servicer, u user.Servicer, logger *log.Logger) *UnshareNoteUseCase {
	return &UnshareNoteUseCase{
		NoteService: s,
		UserService: u,
		logger:      logger,
	}
}
