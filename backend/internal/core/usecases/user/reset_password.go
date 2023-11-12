package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
)

func (s Service) ResetPassword(ctx context.Context, r requests.ResetPasswordRequest) error {
	//match := s.PasswordMatchesRepeatPassword(ctx, r.Password, r.RepeatPassword)
	//if !match {
	//	return errors.New("passwords do not match")
	//}
	//
	//encryptedPassword := s.EncryptPassword(ctx, r.Password)
	//err := s.repository.ResetPassword(ctx, encryptedPassword, r.ResetPasswordCode)
	//if err != nil {
	//	return err
	//}

	return nil
}
