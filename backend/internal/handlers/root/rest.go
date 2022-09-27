package root

import (
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/todo"
	"log"
	"net/http"
	"strings"
)

type RootHandler struct {
	todoService     ports.TodoServicer
	categoryService ports.CategoryServicer
	logger          *log.Logger
}

func (h RootHandler) Root(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.RequestURI, handlers.TODO_STRING):
		th := todo.NewTodoHandler(h.todoService, h.logger)
		th.HandleTodo(w, r)
	case strings.Contains(r.RequestURI, handlers.CATEGORY_STRING):
		ch := category.NewCategoryHandler(h.categoryService, h.logger)
		ch.HandleCategory(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func NewRootHandler(cs ports.CategoryServicer, ts ports.TodoServicer, logger *log.Logger) RootHandler {
	return RootHandler{
		categoryService: cs,
		todoService:     ts,
		logger:          logger,
	}
}
