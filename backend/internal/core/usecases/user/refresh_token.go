package user

import (
	"context"
	"net/http"
)

func (s Service) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	//userID, err := server.RetrieveJWTClaims(r, nil)
	//if err != nil {
	//	return "", errors.New("invalid token")
	//}
	//u, err := s.repository.GetUser(ctx, uint(int(userID)))
	//if err != nil {
	//	return "", errors.New("invalid token")
	//}
	//
	//newClaims := MyCustomClaims{
	//	UserID: u.ID,
	//	Admin:  u.Admin,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		Subject:   u.Email,
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
	//	},
	//}
	//
	//mySigningKey := pkg2.HmacSampleSecret
	//newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	//ss, err := newToken.SignedString(mySigningKey)
	//if err != nil {
	//	return "", err
	//}
	//
	//return ss, nil
	return "", nil
}
