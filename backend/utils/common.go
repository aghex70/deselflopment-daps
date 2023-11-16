package common

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}

//func FilterUsers(users []domain.User) []domain.FilteredUser {
//	filteredUsers := make([]domain.FilteredUser, 0, len(users))
//	for _, user := range users {
//		filteredUsers = append(filteredUsers, FilterUser(user))
//	}
//	return filteredUsers
//}
//
//func FilterUser(user domain.User) domain.FilteredUser {
//	return domain.FilteredUser{
//		ID:               user.ID,
//		Email:            user.Email,
//		Name:             user.Name,
//		RegistrationDate: user.RegistrationDate,
//	}
//}
