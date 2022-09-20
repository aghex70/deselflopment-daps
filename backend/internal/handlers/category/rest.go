package category

import (
	"github.com/aghex70/daps/internal/core/ports"
	"log"
	"net/http"
)

type CategoryHandler struct {
	categoryService ports.CategoryServicer
	logger          *log.Logger
}

func (h CategoryHandler) Category(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.DeleteCategory(w, r)
	case http.MethodGet:
		h.GetCategory(w, r)
	case http.MethodPost:
		h.CreateCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func NewCategoryHandler(cs ports.CategoryServicer, logger *log.Logger) CategoryHandler {
	return CategoryHandler{
		categoryService: cs,
		logger:          logger,
	}
}
