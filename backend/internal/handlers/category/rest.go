package category

import (
	"encoding/json"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
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

func (h CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) error {
	payload := ports.CreateCategoryRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return err
	}

	err = h.categoryService.Create(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return err
	}
	return nil
}

func (h CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.List(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.ListCategoriesResponse{Categories: categories})
	w.Write(b)
}

func NewCategoryHandler(cs ports.CategoryServicer, logger *log.Logger) CategoryHandler {
	return CategoryHandler{
		categoryService: cs,
		logger:          logger,
	}
}
