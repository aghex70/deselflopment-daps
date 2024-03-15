package note

import "github.com/aghex70/daps/internal/ports/domain"

type ListNotesResponse struct {
	Notes []domain.Note `json:"notes"`
}
