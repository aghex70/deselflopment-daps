package user

import (
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type UserHandler struct {
	userService ports.UserServicer
	logger      *log.Logger
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	payload := ports.CreateUserRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, w)
		return
	}

	err = h.userService.Register(nil, payload)
	if err != nil {
		handlers.ThrowError(err, w)
		return
	}
}

func (h UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func NewUserHandler(us ports.UserServicer, logger *log.Logger) UserHandler {
	return UserHandler{
		userService: us,
		logger:      logger,
	}
}
