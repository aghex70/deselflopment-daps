package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	categoryService ports.CategoryServicer
	toDoService     ports.TodoServicer
	logger          *log.Logger
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
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
	payload := ports.CreateCategoryRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Create(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request, id int) {
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	////w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	////w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//
	//if r.Method == "OPTIONS" {
	//	http.Error(w, "No Content", http.StatusNoContent)
	//	return
	//}
	payload := ports.UpdateCategoryRequest{CategoryId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Update(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}
	payload := ports.DeleteCategoryRequest{CategoryId: int64(id)}
	fmt.Println("cccccccccccccccccccccccccccc11111111111111111111111")
	err := handlers.ValidateRequest(r, &payload)
	fmt.Println("cccccccccccccccccccccccccccc22222222222222222222222")
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.categoryService.Delete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request, id int) {
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	////w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	////w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//
	//if r.Method == "OPTIONS" {
	//	http.Error(w, "No Content", http.StatusNoContent)
	//	return
	//}
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@")
	fmt.Println("11111111111111111111")

	payload := ports.GetCategoryRequest{CategoryId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	fmt.Println("22222222222222222222")
	if err != nil {
		fmt.Println("33333333333333333333")
		fmt.Printf("err: %+v", err)
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	c, err := h.categoryService.Get(nil, r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(c)
	w.Write(b)
}

func (h CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

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
