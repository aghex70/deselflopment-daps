package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
)

func (s Service) Activate(ctx context.Context, r requests.ActivateUserRequest) error {
	//err := s.repository.ActivateUser(ctx, r.ActivationCode)
	//if err != nil {
	//	return err
	//}

	return nil
}
