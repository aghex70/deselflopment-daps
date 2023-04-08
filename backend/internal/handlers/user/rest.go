package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/pkg"
	"gorm.io/gorm"
)

type Handler struct {
	userService ports.UserServicer
}

func (h Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
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

	err = h.userService.Register(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
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

	token, userId, err := h.userService.Login(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token, UserId: userId})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	token, err := h.userService.RefreshToken(context.TODO(), r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	_, err = h.userService.CheckAdmin(context.TODO(), r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.DeleteUserRequest{UserId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Delete(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.GetUserRequest{UserId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	user, err := h.userService.Get(context.TODO(), r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	filteredUser := pkg.FilterUser(user)
	b, err := json.Marshal(filteredUser)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
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

	err = h.userService.ProvisionDemoUser(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := handlers.CheckHttpMethod(http.MethodGet, w, r)
	if err != nil {
		return
	}

	users, err := h.userService.List(context.TODO(), r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	filteredUsers := pkg.FilterUsers(users)
	b, err := json.Marshal(handlers.ListUsersResponse{Users: filteredUsers})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) ImportCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Accept", "multipart/form-data")
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	// Parse the CSV file from the request
	f, _, err := r.FormFile("todos.csv")
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	defer f.Close()

	err = h.userService.ImportCSV(context.TODO(), r, f)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ActivateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	payload := ports.ActivateUserRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.Activate(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) ResetLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	payload := ports.ResetLinkRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.SendResetLink(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	payload := ports.ResetPasswordRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.ResetPassword(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewUserHandler(us ports.UserServicer) Handler {
	return Handler{
		userService: us,
	}
}
