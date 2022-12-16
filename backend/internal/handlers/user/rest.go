package user

import (
	"encoding/json"
	"errors"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/pkg"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UserHandler struct {
	userService ports.UserServicer
	logger      *log.Logger
}

func (h UserHandler) HandleUser(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.USER_STRING)[1]
	userId, err := strconv.Atoi(path)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetUser(w, r, userId)
	case http.MethodDelete:
		h.DeleteUser(w, r, userId)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3100")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	}
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
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3100")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	}
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
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3100")
	} else {
		w.Header().Add("Access-Control-Allow-Origin", "http://deselflopment.com")
	}
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

func (h UserHandler) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	err = h.userService.CheckAdmin(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.DeleteUserRequest{UserId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Delete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.GetUserRequest{UserId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	user, err := h.userService.Get(nil, r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(user)
	w.Write(b)
}

func (h UserHandler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
	payload := ports.ProvisionDemoUserRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	err = h.userService.ProvisionDemoUser(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := handlers.CheckHttpMethod(http.MethodGet, w, r)
	if err != nil {
		return
	}

	users, err := h.userService.List(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	filteredUsers := pkg.FilterUsers(users)
	b, err := json.Marshal(handlers.ListUsersResponse{Users: filteredUsers})
	w.Write(b)
}

func NewUserHandler(us ports.UserServicer, logger *log.Logger) UserHandler {
	return UserHandler{
		userService: us,
		logger:      logger,
	}
}
