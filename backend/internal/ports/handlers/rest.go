package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"reflect"
	"strings"
)

const (
	CATEGORY_STRING    string = "categories/"
	TODO_STRING        string = "todos/"
	TOPIC_STRING       string = "topics/"
	NOTE_STRING        string = "notes/"
	SHARE_STRING       string = "/share"
	UNSHARE_STRING     string = "/unshare"
	UNSUBSCRIBE_STRING string = "/unsubscribe"
	START_STRING       string = "/start"
	COMPLETE_STRING    string = "/complete"
	RESTART_STRING     string = "/restart"
	ACTIVATE_STRING    string = "/activate"
	USER_STRING        string = "users/"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

func ValidateRequest(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil && err != io.EOF {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		return err
	}

	return nil
}

func ThrowError(err error, status int, w http.ResponseWriter) {
	resp := APIErrorResponse{
		Message: err.Error(),
	}
	data, _ := json.Marshal(resp)
	w.WriteHeader(status)
	_, werr := w.Write(data)
	if werr != nil {
		return
	}
}

func CheckHttpMethod(status string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != status {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return pkg.MethodNotAllowedError
	}
	return nil
}

func RetrieveHeaderJWT(r *http.Request) string {
	authorizationHeader := r.Header["Authorization"][0]
	headerToken := strings.Split(authorizationHeader, " ")[1]
	return headerToken
}

func RetrieveBodyJWT(payload interface{}) string {
	value := reflect.ValueOf(payload)
	bodyToken := value.FieldByName("AccessToken").String()
	return bodyToken
}

func RetrieveJWTClaims(r *http.Request, payload interface{}) (uint, error) {
	var tokenString string
	if r.Header["Authorization"] != nil {
		tokenString = RetrieveHeaderJWT(r)
	} else {
		tokenString = RetrieveBodyJWT(payload)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return pkg.HmacSampleSecret, nil
	})
	if err != nil {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	return userID, nil
}

func CheckJWTClaims(userID, JWTUserID uint) error {
	if userID != JWTUserID {
		return pkg.UnauthorizedError
	}
	return nil
}

func JWTAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Header["Authorization"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorizationHeader := r.Header["Authorization"][0]
		headerToken := strings.Split(authorizationHeader, " ")[1]
		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return pkg.HmacSampleSecret, nil
		})

		if err != nil {
			ThrowError(err, http.StatusUnauthorized, w)
			return
		}
		if !token.Valid {
			ThrowError(err, http.StatusUnauthorized, w)
			return
		}
		f(w, r)
	}
}
