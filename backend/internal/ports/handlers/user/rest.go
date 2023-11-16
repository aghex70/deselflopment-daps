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
	CheckAdminUseCase        user.CheckAdminUseCase
	DeleteUserUseCase        user.DeleteUserUseCase
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
	//case http.MethodGet:
	//	h.GetUser(w, r, userID)
	case http.MethodDelete:
		h.DeleteUser(w, r, userID)
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
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.RegisterUserUseCase.Execute(context.TODO(), payload)
	if err != nil {
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
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := requests.LoginUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
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

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := requests.RefreshTokenRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	// Check if user_id from the payload matches the user_id from the token
	err = handlers.CheckJWTClaims(userID, payload.UserID)
	if err != nil {
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

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	payload := requests.ResetLinkRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.SendResetLinkUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	payload := requests.ResetPasswordRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.ResetPasswordUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Activate(w http.ResponseWriter, r *http.Request) {
	pkg.SetCORSHeaders(w, r)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusMethodNotAllowed, w)
		return
	}

	payload := requests.ActivateUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.ActivateUserUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	_, err = h.CheckAdminUseCase.Execute(context.TODO(), userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	//return
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	//	payload := requests.DeleteUserRequest{UserID: uint(int64(id))}
	//	err := handlers.ValidateRequest(r, &payload)
	//	if err != nil {
	//		handlers.ThrowError(err, http.StatusBadRequest, w)
	//		return
	//	}
	//
	//	err = h.userService.Delete(context.TODO(), r, payload)
	//	if err != nil {
	//		handlers.ThrowError(err, http.StatusBadRequest, w)
	//		return
	//	}
	//	w.WriteHeader(http.StatusNoContent)
	//}
	//
	//func (h Handler) GetUser(w http.ResponseWriter, r *http.Request, id int) {
	//	payload := requests.GetUserRequest{UserID: uint(int64(id))}
	//	err := handlers.ValidateRequest(r, &payload)
	//	if err != nil {
	//		handlers.ThrowError(err, http.StatusBadRequest, w)
	//		return
	//	}
	//
	//	//user, err := h.userService.Get(context.TODO(), r, payload)
	//	_, err = h.userService.Get(context.TODO(), r, payload)
	//	if err != nil {
	//		if errors.Is(err, gorm.ErrRecordNotFound) {
	//			w.WriteHeader(http.StatusNotFound)
	//			return
	//		}
	//		handlers.ThrowError(err, http.StatusBadRequest, w)
	//		return
	//	}
	//	return
	//	//filteredUser := pkg.FilterUser(user)
	//	//b, err := json.Marshal(filteredUser)
	//	//if err != nil {
	//	//	return
	//	//}
	//	//_, err = w.Write(b)
	//	//if err != nil {
	//	//	return
	//	//}
}

func (h Handler) ProvisionDemoUser(w http.ResponseWriter, r *http.Request) {
	//payload := requests.ProvisionDemoUserRequest{}
	//err := handlers.ValidateRequest(r, &payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//
	//err = handlers.CheckHttpMethod(http.MethodPost, w, r)
	//if err != nil {
	//	return
	//}
	//
	//err = h.userService.ProvisionDemoUser(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//w.WriteHeader(http.StatusCreated)
}

func (h Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	//err := handlers.CheckHttpMethod(http.MethodGet, w, r)
	//if err != nil {
	//	return
	//}
	//
	////users, err := h.userService.List(context.TODO(), r)
	//_, err = h.userService.List(context.TODO(), r)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//return
	////filteredUsers := pkg.FilterUsers(users)
	////b, err := json.Marshal(handlers.ListUsersResponse{Users: filteredUsers})
	////if err != nil {
	////	return
	////}
	////_, err = w.Write(b)
	////if err != nil {
	////	return
	////}
}

func (h Handler) ImportCSV(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Accept", "multipart/form-data")
	//err := handlers.CheckHttpMethod(http.MethodPost, w, r)
	//if err != nil {
	//	return
	//}
	//
	//// Parse the CSV file from the request
	//f, _, err := r.FormFile("todos.csv")
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//defer f.Close()
	//
	//err = h.userService.ImportCSV(context.TODO(), r, f)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//w.WriteHeader(http.StatusCreated)
}

func NewUserHandler(
	activateUserUseCase *user.ActivateUserUseCase,
	//checkAdminUseCase *user.CheckAdminUseCase,
	//deleteUserUseCase *user.DeleteUserUseCase,
	//getUserUseCase *user.GetUserUseCase,
	//listUsersUseCase *user.ListUsersUseCase,
	loginUserUseCase *user.LoginUserUseCase,
	//provisionDemoUserUseCase *user.ProvisionDemoUserUseCase,
	refreshTokenUseCase *user.RefreshTokenUseCase,
	registerUserUseCase *user.RegisterUserUseCase,
	resetPasswordUseCase *user.ResetPasswordUseCase,
	sendResetLinkUseCase *user.SendResetLinkUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		ActivateUserUseCase: *activateUserUseCase,
		//CheckAdminUseCase:        checkAdminUseCase,
		//DeleteUserUseCase:        deleteUserUseCase,
		//GetUserUseCase:           getUserUseCase,
		//ListUsersUseCase:         listUsersUseCase,
		LoginUserUseCase: *loginUserUseCase,
		//ProvisionDemoUserUseCase: provisionDemoUserUseCase,
		RefreshTokenUseCase:  *refreshTokenUseCase,
		RegisterUserUseCase:  *registerUserUseCase,
		ResetPasswordUseCase: *resetPasswordUseCase,
		SendResetLinkUseCase: *sendResetLinkUseCase,
		logger:               logger,
	}
}
