package user

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/handlers"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	ActivateUserUseCase      user.ActivateUserUseCase
	DeleteUserUseCase        user.DeleteUserUseCase
	GetUserUseCase           user.GetUserUseCase
	ListUsersUseCase         user.ListUsersUseCase
	LoginUserUseCase         user.LoginUserUseCase
	ProvisionDemoUserUseCase user.ProvisionDemoUserUseCase
	RefreshTokenUseCase      user.RefreshTokenUseCase
	RegisterUserUseCase      user.RegisterUserUseCase
	ResetPasswordUseCase     user.ResetPasswordUseCase
	SendResetLinkUseCase     user.SendResetLinkUseCase
	UpdateUserUseCase        user.UpdateUserUseCase
	logger                   *log.Logger
}

func (h Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	path := strings.Split(r.RequestURI, handlers.USER_STRING)[1]
	userID, err := strconv.Atoi(path)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.Get(w, r, uint(userID))
	case http.MethodDelete:
		h.Delete(w, r, uint(userID))
	case http.MethodPut:
		h.Update(w, r, uint(userID))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := requests.CreateUserRequest{}
	if err = handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	if err = h.RegisterUserUseCase.Execute(context.TODO(), payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.LoginUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, userID, err := h.LoginUserUseCase.Execute(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token, UserID: userID})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.RefreshTokenRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	// Check if user_id from the payload matches the user_id from the token
	if err = handlers.CheckJWTClaims(userID, payload.UserID); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		return
	}
	token, _, err := h.RefreshTokenUseCase.Execute(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token, UserID: userID})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) ResetLink(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.ResetLinkRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err := h.SendResetLinkUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.ResetPasswordRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.ResetPasswordUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Activate(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.ActivateUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.ActivateUserUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	payload.UserID = id
	if err = h.UpdateUserUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodGet, w, r); err != nil {
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	users, err := h.ListUsersUseCase.Execute(context.TODO(), nil, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	//return
	//filteredUsers := pkg.FilterUsers(users)
	//b, err := json.Marshal(handlers.ListUsersResponse{Users: filteredUsers})
	b, err := json.Marshal(handlers.ListUsersResponse{Users: users})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if err := handlers.CheckHttpMethod(http.MethodPost, w, r); err != nil {
		return
	}

	payload := requests.ProvisionDemoUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.ProvisionDemoUserUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.DeleteUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	payload.UserID = id
	if err = h.DeleteUserUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetUserRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	payload.UserID = id
	u, err := h.GetUserUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.GetUserResponse{User: u})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func NewUserHandler(
	activateUserUseCase *user.ActivateUserUseCase,
	deleteUserUseCase *user.DeleteUserUseCase,
	getUserUseCase *user.GetUserUseCase,
	listUsersUseCase *user.ListUsersUseCase,
	loginUserUseCase *user.LoginUserUseCase,
	provisionDemoUserUseCase *user.ProvisionDemoUserUseCase,
	refreshTokenUseCase *user.RefreshTokenUseCase,
	registerUserUseCase *user.RegisterUserUseCase,
	resetPasswordUseCase *user.ResetPasswordUseCase,
	sendResetLinkUseCase *user.SendResetLinkUseCase,
	updateUserUseCase *user.UpdateUserUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		ActivateUserUseCase:      *activateUserUseCase,
		DeleteUserUseCase:        *deleteUserUseCase,
		GetUserUseCase:           *getUserUseCase,
		ListUsersUseCase:         *listUsersUseCase,
		LoginUserUseCase:         *loginUserUseCase,
		ProvisionDemoUserUseCase: *provisionDemoUserUseCase,
		RefreshTokenUseCase:      *refreshTokenUseCase,
		RegisterUserUseCase:      *registerUserUseCase,
		ResetPasswordUseCase:     *resetPasswordUseCase,
		SendResetLinkUseCase:     *sendResetLinkUseCase,
		UpdateUserUseCase:        *updateUserUseCase,
		logger:                   logger,
	}
}
