package category

import (
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/category"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	//categoryService category.Servicer
}

func (h Handler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.CATEGORY_STRING)[1]

	categoryID, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	categoryIDUint := uint(categoryID)

	switch r.Method {
	case http.MethodGet:
		h.GetCategory(w, r, categoryIDUint)
	case http.MethodDelete:
		h.DeleteCategory(w, r, categoryIDUint)
	case http.MethodPut:
		h.UpdateCategory(w, r, categoryIDUint)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	payload := requests.CreateCategoryRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.categoryService.Create(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) UpdateCategory(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateCategoryRequest{CategoryID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.categoryService.Update(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
}

func (h Handler) DeleteCategory(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.DeleteCategoryRequest{CategoryID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.categoryService.Delete(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) GetCategory(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetCategoryRequest{CategoryID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//c, err := h.categoryService.Get(context.TODO(), r, payload)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		w.WriteHeader(http.StatusNotFound)
	//		return
	//	}
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//b, err := json.Marshal(c)
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
	return
}

func (h Handler) ListCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		http.Error(w, "No Content", http.StatusNoContent)
		return
	}

	//categories, err := h.categoryService.List(context.TODO(), r)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//
	//b, err := json.Marshal(handlers.ListCategoriesResponse{Categories: categories})
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
	return
}

//func NewCategoryHandler(cs category.Servicer) Handler {
//	return Handler{
//categoryService: cs,
//}
//}
