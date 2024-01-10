package requests

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=13"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=13"`
}

type RefreshTokenRequest struct {
	UserID uint `json:"user_id"`
}

type UpdateUserRequest struct {
	AutoSuggest *bool   `json:"auto_suggest"`
	Language    *string `json:"language"`
	UserID      uint    `json:"user_id"`
}

type ProvisionDemoUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	UserID uint `json:"user_id"`
}

type GetUserRequest struct {
	UserID uint `json:"user_id"`
}

type ActivateUserRequest struct {
	ActivationCode string `json:"activation_code"`
}

type ResetPasswordRequest struct {
	Password          string `json:"password" validate:"required,min=13"`
	ResetPasswordCode string `json:"reset_password_code"`
}

type ResetLinkRequest struct {
	Email string `json:"email" validate:"required,email"`
}
