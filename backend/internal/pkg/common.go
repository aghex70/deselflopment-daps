package pkg

import (
	"errors"
	"os"
)

var (
	JWTSigningKey        = os.Getenv("JWT_SIGNING_KEY")
	HmacSampleSecret     = []byte(JWTSigningKey)
	BaseCategoriesIds    = []uint{1, 2, 3, 4, 5}
	BaseCategoryIdFilter = map[string]interface{}{"id": 1}
	DemoUserName         = os.Getenv("DEMO_USER_NAME")
	ProjectName          = os.Getenv("PROJECT_NAME")
)

var (
	FromName          = os.Getenv("PROJECT_NAME")
	FromEmail         = os.Getenv("FROM_EMAIL")
	SendGridApiKey    = os.Getenv("SENDGRID_API_KEY")
	DapsLocalUrl      = os.Getenv("DAPS_LOCAL_URL")
	DapsProductionUrl = os.Getenv("DAPS_PRODUCTION_URL")
)
var (
	Environment        = os.Getenv("ENVIRONMENT")
	ActivationCodeLink = GetOrigin() + "/activate/"
	ResetPasswordLink  = GetOrigin() + "/reset-password/"
)

// var DeselflopmentLocalUrl = os.Getenv("DESELFLOPMENT_LOCAL_URL")
// var DeselflopmentProductionUrl = os.Getenv("DESELFLOPMENT_PRODUCTION_URL")

// User
var (
	UserAlreadyRegisteredError = errors.New("user already registered")
	InvalidCredentialsError    = errors.New("invalid credentials")
	UnauthorizedError          = errors.New("unauthorized")
)
