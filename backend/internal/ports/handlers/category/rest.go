package category

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/category"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/responses"
	categoryResponses "github.com/aghex70/daps/internal/ports/responses/category"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	CreateCategoryUseCase      category.CreateCategoryUseCase
	DeleteCategoryUseCase      category.DeleteCategoryUseCase
	GetCategoryUseCase         category.GetCategoryUseCase
	GetSummaryUseCase          category.GetSummaryUseCase
	ListCategoriesUseCase      category.ListCategoriesUseCase
	ListCategoryUsersUseCase   category.ListCategoryUsersUseCase
	ShareCategoryUseCase       category.ShareCategoryUseCase
	UnshareCategoryUseCase     category.UnshareCategoryUseCase
	UnsubscribeCategoryUseCase category.UnsubscribeCategoryUseCase
	UpdateCategoryUseCase      category.UpdateCategoryUseCase
	logger                     *log.Logger
}

func (h Handler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.List(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	payload := requests.CreateCategoryRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	c, err := h.CreateCategoryUseCase.Execute(context.TODO(), userID, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(responses.CreateEntityResponse{ID: c.ID})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	categories, err := h.ListCategoriesUseCase.Execute(context.TODO(), nil, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	filteredCategories := pkg.FilterCategories(categories)
	b, err := json.Marshal(categoryResponses.ListCategoriesResponse{Categories: filteredCategories})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	// Get category id & action (if present) from request URI
	path := strings.Split(r.RequestURI, handlers.CATEGORY_STRING)[1]
	c := strings.Split(path, "/")[0]
	categoryID, err := strconv.Atoi(c)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	if strings.Contains(r.RequestURI, handlers.SHARE_STRING) {
		h.Share(w, r, uint(categoryID))
	} else if strings.Contains(r.RequestURI, handlers.UNSHARE_STRING) {
		h.Unshare(w, r, uint(categoryID))
	} else if strings.Contains(r.RequestURI, handlers.UNSUBSCRIBE_STRING) {
		h.Unsubscribe(w, r, uint(categoryID))
	} else if strings.Contains(r.RequestURI, handlers.USER_STRING) {
		h.ListCategoryUsers(w, r, uint(categoryID))
	} else {
		switch r.Method {
		case http.MethodGet:
			h.Get(w, r, uint(categoryID))
		case http.MethodDelete:
			h.Delete(w, r, uint(categoryID))
		case http.MethodPut:
			h.Update(w, r, uint(categoryID))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	c, err := h.GetCategoryUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	filteredCategory := pkg.FilterCategory(c)
	b, err := json.Marshal(filteredCategory)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.DeleteCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.DeleteCategoryUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UpdateCategoryUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Share(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.ShareCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.ShareCategoryUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Unshare(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UnshareCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UnshareCategoryUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Unsubscribe(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UnsubscribeCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UnsubscribeCategoryUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetSummary(w http.ResponseWriter, r *http.Request) {
	if err := handlers.CheckHttpMethod(http.MethodGet, w, r); err != nil {
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	summary, err := h.GetSummaryUseCase.Execute(context.TODO(), userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(summary)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h Handler) ListCategoryUsers(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetCategoryRequest{CategoryID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	if err := handlers.CheckHttpMethod(http.MethodGet, w, r); err != nil {
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	users, err := h.ListCategoryUsersUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(categoryResponses.ListCategoryUsersResponse{Users: users})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func NewCategoryHandler(
	createCategoryUseCase *category.CreateCategoryUseCase,
	deleteCategoryUseCase *category.DeleteCategoryUseCase,
	getCategoryUseCase *category.GetCategoryUseCase,
	getSummaryUseCase *category.GetSummaryUseCase,
	listCategoriesUseCase *category.ListCategoriesUseCase,
	listCategoryUsersUseCase *category.ListCategoryUsersUseCase,
	shareCategoryUseCase *category.ShareCategoryUseCase,
	unshareCategoryUseCase *category.UnshareCategoryUseCase,
	unsubscribeCategoryUseCase *category.UnsubscribeCategoryUseCase,
	updateCategoryUseCase *category.UpdateCategoryUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		CreateCategoryUseCase:      *createCategoryUseCase,
		DeleteCategoryUseCase:      *deleteCategoryUseCase,
		GetCategoryUseCase:         *getCategoryUseCase,
		GetSummaryUseCase:          *getSummaryUseCase,
		ListCategoriesUseCase:      *listCategoriesUseCase,
		ListCategoryUsersUseCase:   *listCategoryUsersUseCase,
		ShareCategoryUseCase:       *shareCategoryUseCase,
		UnshareCategoryUseCase:     *unshareCategoryUseCase,
		UnsubscribeCategoryUseCase: *unsubscribeCategoryUseCase,
		UpdateCategoryUseCase:      *updateCategoryUseCase,
		logger:                     logger,
	}
}
