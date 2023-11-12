package user

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"net/http"
)

func (s Service) Get(ctx context.Context, r *http.Request, req requests.GetUserRequest) (domain2.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return domain2.User{}, err
	//}
	//
	//u, err := s.repository.GetUser(ctx, uint(int(req.UserID)))
	//if err != nil {
	//	return domain2.User{}, err
	//}
	//
	//return u, nil
	return domain2.User{}, nil
}
