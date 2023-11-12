package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
)

func (s Service) SendResetLink(ctx context.Context, r requests.ResetLinkRequest) error {
	//u, err := s.repository.CreateResetLink(ctx, r.Email)
	//if err != nil {
	//	return err
	//}
	//
	//e := domain.Email{
	//	From:      pkg.FromEmail,
	//	To:        u.Email,
	//	Recipient: u.Name,
	//	Subject:   "📣 DAPS - Password reset request 📣",
	//	Body:      "In order to reset your password, please follow this link: " + pkg.ResetPasswordLink + u.ResetPasswordCode,
	//	User:      u.ID,
	//}
	//
	//err = pkg.SendEmail(e)
	//if err != nil {
	//	fmt.Printf("Error sending email: %+v", err)
	//	e.Error = err.Error()
	//	e.Sent = false
	//	_, errz := s.repository.CreateEmail(ctx, e)
	//	if errz != nil {
	//		fmt.Printf("Error saving email: %+v", errz)
	//		return errz
	//	}
	//	return err
	//}
	//
	//e.Sent = true
	//_, err = s.repository.CreateEmail(ctx, e)
	//if err != nil {
	//	return err
	//}
	return nil
}
