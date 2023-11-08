package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aghex70/daps/internal/common/pkg"
	handlers2 "github.com/aghex70/daps/internal/ports/handlers"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/user"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Handler struct {
	userService user.Servicer
}

func (h Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers2.USER_STRING)[1]
	userID, err := strconv.Atoi(path)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetUser(w, r, userID)
	case http.MethodDelete:
		h.DeleteUser(w, r, userID)
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

	err := handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := requests.CreateUserRequest{}
	err = handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Register(context.TODO(), payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
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
	err := handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := requests.LoginUserRequest{}
	err = handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, userID, err := h.userService.Login(context.TODO(), payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers2.TokenResponse{AccessToken: token, UserID: userID})
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
	err := handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	token, err := h.userService.RefreshToken(context.TODO(), r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers2.TokenResponse{AccessToken: token})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	err := handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	_, err = h.userService.CheckAdmin(context.TODO(), r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := requests.DeleteUserRequest{UserID: uint(int64(id))}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Delete(context.TODO(), r, payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	payload := requests.GetUserRequest{UserID: uint(int64(id))}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//user, err := h.userService.Get(context.TODO(), r, payload)
	_, err = h.userService.Get(context.TODO(), r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	return
	//filteredUser := pkg.FilterUser(user)
	//b, err := json.Marshal(filteredUser)
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
}

func (h Handler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
	payload := requests.ProvisionDemoUserRequest{}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	err = h.userService.ProvisionDemoUser(context.TODO(), r, payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	err := handlers2.CheckHttpMethod(http.MethodGet, w, r)
	if err != nil {
		return
	}

	//users, err := h.userService.List(context.TODO(), r)
	_, err = h.userService.List(context.TODO(), r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	return
	//filteredUsers := pkg.FilterUsers(users)
	//b, err := json.Marshal(handlers.ListUsersResponse{Users: filteredUsers})
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
}

func (h Handler) ImportCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Accept", "multipart/form-data")
	err := handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	// Parse the CSV file from the request
	f, _, err := r.FormFile("todos.csv")
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	defer f.Close()

	err = h.userService.ImportCSV(context.TODO(), r, f)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
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

	payload := requests.ActivateUserRequest{}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.Activate(context.TODO(), payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
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

	payload := requests.ResetLinkRequest{}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.SendResetLink(context.TODO(), payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
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

	payload := requests.ResetPasswordRequest{}
	err := handlers2.ValidateRequest(r, &payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = handlers2.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers2.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	err = h.userService.ResetPassword(context.TODO(), payload)
	if err != nil {
		handlers2.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewUserHandler(us user.Servicer) Handler {
	return Handler{
		userService: us,
	}
}
