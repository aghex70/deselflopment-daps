package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"net/http"
)

func (s Service) Delete(ctx context.Context, r *http.Request, req requests.DeleteUserRequest) error {
	////adminID, err := s.CheckAdmin(ctx, r)
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//err = s.repository.DeleteUser(ctx, uint(int(req.UserID)))
	//if err != nil {
	//	return err
	//}

	return nil
}
