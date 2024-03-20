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
	"sync"
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
			return domain.Note{}, err
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

	n := domain.Note{
		Content: r.Content,
		Topics:  topics,
		Shared:  false,
		OwnerID: u.ID,
		Users:   []domain.User{u},
	}
	nn, err := uc.NoteService.Create(ctx, n)
	if err != nil {
		return domain.Note{}, err
	}

	return nn, nil
}

func NewCreateNoteUseCase(s note.Servicer, u user.Servicer, t topic.Servicer, logger *log.Logger) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		NoteService:  s,
		UserService:  u,
		TopicService: t,
		logger:       logger,
	}
}
