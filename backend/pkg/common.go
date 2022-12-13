package pkg

import "os"

var JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")
var HmacSampleSecret = []byte(JWTSigningKey)
var BaseCategoriesIds = []int{1, 2, 3, 4, 5}
var DemoUserPassword = os.Getenv("DEMO_USER_PASSWORD")
var DemoUserName = os.Getenv("DEMO_USER_NAME")
