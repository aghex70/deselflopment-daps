package pkg

import (
	"fmt"
	"os"
)

var JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")
var HmacSampleSecret = []byte(JWTSigningKey)
var BaseCategoriesIds = []int{1, 2, 3, 4, 5}
var DemoUserName = os.Getenv("DEMO_USER_NAME")
var MaximumConcurrentSuggestions = 4
var FromName = os.Getenv("FROM_NAME")
var FromEmail = os.Getenv("FROM_EMAIL")
var SendGridApiKey = os.Getenv("SENDGRID_API_KEY")
var FrontendPort = os.Getenv("FRONTEND_PORT")
var DevelopUrl = fmt.Sprintf("http://localhost:%s", FrontendPort)
var ProductionUrl = os.Getenv("PRODUCTION_URL")
var Environment = os.Getenv("ENVIRONMENT")
var ActivationCodeLink = GetOrigin() + "/activate/"
var ResetPasswordLink = GetOrigin() + "/reset-password/"
