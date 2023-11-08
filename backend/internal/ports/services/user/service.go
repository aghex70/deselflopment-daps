package user

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"mime/multipart"
	"net/http"
)

type Servicer interface {
	Register(ctx context.Context, r requests.CreateUserRequest) error
	Login(ctx context.Context, r requests.LoginUserRequest) (string, int, error)
	RefreshToken(ctx context.Context, r *http.Request) (string, error)
	CheckAdmin(ctx context.Context, r *http.Request) (int, error)
	Delete(ctx context.Context, r *http.Request, req requests.DeleteUserRequest) error
	Get(ctx context.Context, r *http.Request, req requests.GetUserRequest) (domain2.User, error)
	ProvisionDemoUser(ctx context.Context, r *http.Request, req requests.ProvisionDemoUserRequest) error
	List(ctx context.Context, r *http.Request) ([]domain2.User, error)
	ImportCSV(ctx context.Context, r *http.Request, f multipart.File) error
	Activate(ctx context.Context, r requests.ActivateUserRequest) error
	SendResetLink(ctx context.Context, r requests.ResetLinkRequest) error
	ResetPassword(ctx context.Context, r requests.ResetPasswordRequest) error
}
