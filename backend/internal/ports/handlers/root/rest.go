package root

import (
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/handlers/category"
	"github.com/aghex70/daps/internal/ports/handlers/todo"
	"github.com/aghex70/daps/internal/ports/handlers/user"
	category2 "github.com/aghex70/daps/internal/ports/services/category"
	todo2 "github.com/aghex70/daps/internal/ports/services/todo"
	user2 "github.com/aghex70/daps/internal/ports/services/user"
	"net/http"
	"strings"
)

type Handler struct {
	todoService     todo2.Servicer
	categoryService category2.Servicer
	userService     user2.Servicer
}

func (h Handler) Root(w http.ResponseWriter, r *http.Request) {
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

func NewRootHandler(cs category2.Servicer, ts todo2.Servicer, us user2.Servicer) Handler {
	return Handler{
		categoryService: cs,
		todoService:     ts,
		userService:     us,
	}
}
