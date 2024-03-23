package note

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/services/note"
	"github.com/aghex70/daps/internal/ports/services/topic"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/note"
	"sync"

	"log"
)

type UpdateNoteUseCase struct {
	NoteService  note.Servicer
	TopicService topic.Servicer
	UserService  user.Servicer
	logger       *log.Logger
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

	// Create a wait group
	var wg sync.WaitGroup

	// Create a channel to receive errors
	ec := make(chan error, len(r.TopicIDs))

	// Create a channel to receive topics
	tc := make(chan domain.Topic, len(r.TopicIDs))

	// Create a goroutine for each topic ID
	for _, id := range r.TopicIDs {
		// Increment the wait group counter
		wg.Add(1)

		// Create a goroutine
		go func(id uint) {
			// Decrement the wait group counter when the goroutine finishes
			defer wg.Done()

			// Fetch the topic for the current ID
			t, err := uc.TopicService.Get(ctx, id)
			if err != nil {
				// Send the error to the error channel
				ec <- err
				return
			}
			tc <- t
		}(id)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the error channel to signal that no more errors will be sent
	close(ec)

	// Close the topic channel to signal that no more topics will be sent
	close(tc)

	// Check if there were any errors
	select {
	case err := <-ec:
		if err != nil {
			return err
		}
	default:
		// No error, continue
	}

	// Create a slice to store fetched topics
	var topics []domain.Topic

	// Read and append the fetched topics from the channel to the slice
	for t := range tc {
		topics = append(topics, t)
	}

	un := domain.Note{
		Content: r.Content,
		Topics:  topics,
		Shared:  false,
		OwnerID: u.ID,
		Users:   &[]domain.User{u},
	}

	if err = uc.NoteService.Update(ctx, r.NoteID, un); err != nil {
		return err
	}
	return nil
}

func NewUpdateNoteUseCase(s note.Servicer, t topic.Servicer, u user.Servicer, logger *log.Logger) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{
		NoteService:  s,
		TopicService: t,
		UserService:  u,
		logger:       logger,
	}
}
