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
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
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
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	filters := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		if len(v) == 1 {
			// If there is only one value, use it directly
			filters[k] = v[0]
		} else if len(v) > 1 {
			// If there are multiple values, use a slice
			filters[k] = v
		}
	}
	todos, err := h.ListTodosUseCase.Execute(context.TODO(), &filters, userID)
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

	path := strings.Split(r.RequestURI, handlers.TODO_STRING)[1]
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
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
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
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.DeleteTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateTodoRequest{TodoID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UpdateTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Import(w http.ResponseWriter, r *http.Request) {
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		// Handle the error
		handlers.ThrowError(pkg.ParseFileError, http.StatusBadRequest, w)
		return
	}

	if err = h.ImportTodosUseCase.Execute(context.Background(), userID, file); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

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
