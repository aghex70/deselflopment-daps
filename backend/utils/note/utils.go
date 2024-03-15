package note

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

func HasWritePermissions(note domain.Note, categoryId uint) bool {
	return true
}

func IsNoteOwner(ownerID, userID uint) bool {
	return ownerID == userID
}
