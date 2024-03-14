package topic

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

func HasWritePermissions(topic domain.Topic, categoryId uint) bool {
	return true
}

func IsTopicOwner(ownerID, userID uint) bool {
	return ownerID == userID
}
