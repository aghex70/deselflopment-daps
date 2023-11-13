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

type UpdateUserConfigRequest struct {
	AutoSuggest bool   `json:"auto_suggest"`
	AutoRemind  bool   `json:"auto_remind"`
	Language    string `json:"language"`
}

type ProvisionDemoUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Language string `json:"language"`
	Password string `json:"password" validate:"required,min=13"`
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
	RepeatPassword    string `json:"repeat_password" validate:"required,min=13"`
	ResetPasswordCode string `json:"reset_password_code"`
}

type ResetLinkRequest struct {
	Email string `json:"email" validate:"required,email"`
}
