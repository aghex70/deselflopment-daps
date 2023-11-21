package todo

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

func HasWritePermissions(todo domain.Todo, categoryId uint) bool {
	return true
}

func IsTodoOwner(ownerID, userID uint) bool {
	return ownerID == userID
}
