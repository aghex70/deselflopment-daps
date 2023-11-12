package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
)

func (s Service) Login(ctx context.Context, r requests.LoginUserRequest) (string, int, error) {
	//u, err := s.repository.GetUserByEmail(ctx, r.Email)
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//decryptedPassword, err := s.DecryptPassword(ctx, u.Password)
	//
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//match := s.PasswordsMatch(ctx, decryptedPassword, r.Password)
	//if !match {
	//	return "", 0, errors.New("invalid credentials")
	//}
	//
	//claims := MyCustomClaims{
	//	UserID: u.ID,
	//	Admin:  u.Admin,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		Subject:   r.Email,
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
	//	},
	//}
	//
	//mySigningKey := pkg2.HmacSampleSecret
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//ss, err := token.SignedString(mySigningKey)
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//return ss, int(u.ID), nil
	return "", 0, nil
}
