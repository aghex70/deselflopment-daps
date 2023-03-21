package root

import (
	"net/http"
	"strings"

	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/todo"
	"github.com/aghex70/daps/internal/handlers/user"
)

type RootHandler struct {
	todoService     ports.TodoServicer
	categoryService ports.CategoryServicer
	userService     ports.UserServicer
}

func (h RootHandler) Root(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.RequestURI, handlers.TODO_STRING):
		th := todo.NewTodoHandler(h.todoService)
		th.HandleTodo(w, r)
	case strings.Contains(r.RequestURI, handlers.CATEGORY_STRING):
		ch := category.NewCategoryHandler(h.categoryService)
		ch.HandleCategory(w, r)
	case strings.Contains(r.RequestURI, handlers.USER_STRING):
		uh := user.NewUserHandler(h.userService)
		uh.HandleUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewRootHandler(cs ports.CategoryServicer, ts ports.TodoServicer, us ports.UserServicer) RootHandler {
	return RootHandler{
		categoryService: cs,
		todoService:     ts,
		userService:     us,
	}
}
