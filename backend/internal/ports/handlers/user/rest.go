package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aghex70/daps/internal/core/usecases/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/handlers"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	ActivateUserUseCase      user.ActivateUserUseCase
	DeleteUserUseCase        user.DeleteUserUseCase
	EditProfileUseCase       user.EditProfileUseCase
	GetUserUseCase           user.GetUserUseCase
	ListUsersUseCase         user.ListUsersUseCase
	LoginUserUseCase         user.LoginUserUseCase
	ProvisionDemoUserUseCase user.ProvisionDemoUserUseCase
	RefreshTokenUseCase      user.RefreshTokenUseCase
	RegisterUserUseCase      user.RegisterUserUseCase
	ResetPasswordUseCase     user.ResetPasswordUseCase
	SendResetLinkUseCase     user.SendResetLinkUseCase
	logger                   *log.Logger
}

func (h Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
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
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetProfile(w, r)
	case http.MethodPut:
		h.UpdateProfile(w, r)
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

	token, userID, admin, err := h.LoginUserUseCase.Execute(context.TODO(), payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token, UserID: userID, Admin: admin})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
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

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
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

	// Hide user data
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

func (h Handler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
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

	// Set the demo password and add offset to email
	payload.Name = pkg.DemoUserName
	payload.Password = pkg.DemoUserPassword
	ms := time.Now().UnixNano() / int64(time.Millisecond)
	payload.Email = fmt.Sprintf("%d%s", ms, pkg.DemoUserEmail)

	du, err := h.ProvisionDemoUserUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	// Hide user data
	filteredUser := pkg.FilterUser(du)
	b, err := json.Marshal(handlers.GetUserResponse{User: filteredUser})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
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

	// Hide user data
	filteredUser := pkg.FilterUser(u)
	b, err := json.Marshal(filteredUser)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	if err := handlers.CheckHttpMethod(http.MethodGet, w, r); err != nil {
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	payload := requests.GetUserRequest{}
	payload.UserID = userID
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	u, err := h.GetUserUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	// Hide user data
	profile := pkg.FilterProfile(u)
	b, err := json.Marshal(profile)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	payload := requests.EditProfileRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.EditProfileUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewUserHandler(
	activateUserUseCase *user.ActivateUserUseCase,
	deleteUserUseCase *user.DeleteUserUseCase,
	editProfileUseCase *user.EditProfileUseCase,
	getUserUseCase *user.GetUserUseCase,
	listUsersUseCase *user.ListUsersUseCase,
	loginUserUseCase *user.LoginUserUseCase,
	provisionDemoUserUseCase *user.ProvisionDemoUserUseCase,
	refreshTokenUseCase *user.RefreshTokenUseCase,
	registerUserUseCase *user.RegisterUserUseCase,
	resetPasswordUseCase *user.ResetPasswordUseCase,
	sendResetLinkUseCase *user.SendResetLinkUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		ActivateUserUseCase:      *activateUserUseCase,
		DeleteUserUseCase:        *deleteUserUseCase,
		EditProfileUseCase:       *editProfileUseCase,
		GetUserUseCase:           *getUserUseCase,
		ListUsersUseCase:         *listUsersUseCase,
		LoginUserUseCase:         *loginUserUseCase,
		ProvisionDemoUserUseCase: *provisionDemoUserUseCase,
		RefreshTokenUseCase:      *refreshTokenUseCase,
		RegisterUserUseCase:      *registerUserUseCase,
		ResetPasswordUseCase:     *resetPasswordUseCase,
		SendResetLinkUseCase:     *sendResetLinkUseCase,
		logger:                   logger,
	}
}
