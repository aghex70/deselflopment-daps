package user

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"net/http"
)

func (s Service) List(ctx context.Context, r *http.Request) ([]domain2.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return []domain2.User{}, err
	//}
	//
	//users, err := s.repository.GetUsers(ctx)
	//if err != nil {
	//	return []domain2.User{}, err
	//}
	//
	//return users, nil
	return []domain2.User{}, nil
}
