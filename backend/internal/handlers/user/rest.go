package user

import (
	"github.com/aghex70/daps/internal/core/ports"
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
	panic("foo")
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
