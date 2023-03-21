package category

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	categoryService ports.CategoryServicer
}

func (h CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.CATEGORY_STRING)[1]

	categoryId, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.GetCategory(w, r, categoryId)
	case http.MethodDelete:
		h.DeleteCategory(w, r, categoryId)
	case http.MethodPut:
		h.UpdateCategory(w, r, categoryId)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	payload := ports.CreateCategoryRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Create(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.UpdateCategoryRequest{CategoryId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Update(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.DeleteCategoryRequest{CategoryId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Delete(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.GetCategoryRequest{CategoryId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	c, err := h.categoryService.Get(context.TODO(), r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

	categories, err := h.categoryService.List(context.TODO(), r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.ListCategoriesResponse{Categories: categories})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func NewCategoryHandler(cs ports.CategoryServicer) CategoryHandler {
	return CategoryHandler{
		categoryService: cs,
	}
}
