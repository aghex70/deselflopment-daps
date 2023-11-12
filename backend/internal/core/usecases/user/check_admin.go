package user

import (
	"context"
	"net/http"
)

func (s Service) CheckAdmin(ctx context.Context, r *http.Request) (int, error) {
	//userID, err := server.RetrieveJWTClaims(r, nil)
	//if err != nil {
	//	return 0, errors.New("invalid token")
	//}
	//u, err := s.repository.GetUser(ctx, uint(int(userID)))
	//if err != nil {
	//	return 0, errors.New("invalid token")
	//}
	//
	//if !u.Admin {
	//	return 0, errors.New("unauthorized")
	//}
	//
	//return int(userID), nil
	return 0, nil
}
