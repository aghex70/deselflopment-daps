package user

import "github.com/aghex70/daps/internal/ports/domain"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Admin       bool   `json:"admin"`
	UserID      uint   `json:"user_id"`
}

type GetUserResponse struct {
	User domain.FilteredUser `json:"user"`
}

type ListUsersResponse struct {
	Users []domain.FilteredUser `json:"users"`
}
