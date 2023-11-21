package todo

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/todo"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	CreateTodoUseCase  todo.CreateTodoUseCase
	DeleteTodoUseCase  todo.DeleteTodoUseCase
	GetTodoUseCase     todo.GetTodoUseCase
	ImportTodosUseCase todo.ImportTodosUseCase
	ListTodosUseCase   todo.ListTodosUseCase
	UpdateTodoUseCase  todo.UpdateTodoUseCase
	logger             *log.Logger
}

func (h Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

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
	payload := requests.CreateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	c, err := h.CreateTodoUseCase.Execute(context.TODO(), userID, payload)
	if err != nil {
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
	w.WriteHeader(http.StatusCreated)
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	//q := r.URL.Query()
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	todos, err := h.ListTodosUseCase.Execute(context.TODO(), nil, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.ListTodosResponse{Todos: todos})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	path := strings.Split(r.RequestURI, handlers.CATEGORY_STRING)[1]
	categoryID, err := strconv.Atoi(path)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

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
	//queryParams := strings.Split(path, "?")
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetTodoRequest{TodoID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	t, err := h.GetTodoUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(t)
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
	payload := requests.DeleteTodoRequest{TodoID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.DeleteTodoUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateTodoRequest{TodoID: id}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.UpdateTodoUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Import(w http.ResponseWriter, r *http.Request) {
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		// Handle the error
		handlers.ThrowError(pkg.ParseFileError, http.StatusBadRequest, w)
		return
	}

	err = h.ImportTodosUseCase.Execute(context.Background(), userID, file)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//func (h Handler) ListRecurringTodos(w http.ResponseWriter, r *http.Request) {
//func (h Handler) ListSuggestedTodos(w http.ResponseWriter, r *http.Request) {
//func (h Handler) ListCompletedTodos(w http.ResponseWriter, r *http.Request) {
//func (h Handler) SuggestTodos(w http.ResponseWriter, r *http.Request) {
//func (h Handler) Summary(w http.ResponseWriter, r *http.Request) {

func NewTodoHandler(
	createTodoUseCase *todo.CreateTodoUseCase,
	deleteTodoUseCase *todo.DeleteTodoUseCase,
	getTodoUseCase *todo.GetTodoUseCase,
	importTodosUseCase *todo.ImportTodosUseCase,
	listTodosUseCase *todo.ListTodosUseCase,
	updateTodoUseCase *todo.UpdateTodoUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		CreateTodoUseCase:  *createTodoUseCase,
		DeleteTodoUseCase:  *deleteTodoUseCase,
		GetTodoUseCase:     *getTodoUseCase,
		ImportTodosUseCase: *importTodosUseCase,
		ListTodosUseCase:   *listTodosUseCase,
		UpdateTodoUseCase:  *updateTodoUseCase,
		logger:             logger,
	}
}
