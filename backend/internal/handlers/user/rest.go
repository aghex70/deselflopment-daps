package user

import (
	"encoding/json"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type UserHandler struct {
	userService ports.UserServicer
	logger      *log.Logger
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	////w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	////w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := ports.CreateUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Register(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := ports.LoginUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, userId, err := h.userService.Login(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token, UserId: userId})
	w.Write(b)
}

func (h UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	token, err := h.userService.RefreshToken(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token})
	w.Write(b)
}

func (h UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://3.75.160.227")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	err := handlers.CheckHttpMethod(http.MethodDelete, w, r)
	if err != nil {
		return
	}

	err = h.userService.Remove(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func NewUserHandler(us ports.UserServicer, logger *log.Logger) UserHandler {
	return UserHandler{
		userService: us,
		logger:      logger,
	}
}
