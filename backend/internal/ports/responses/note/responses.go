package note

import "github.com/aghex70/daps/internal/ports/domain"

type ListNotesResponse struct {
	Notes []domain.FilteredNote `json:"notes"`
}

type GetNoteResponse struct {
	domain.FilteredNote
}
