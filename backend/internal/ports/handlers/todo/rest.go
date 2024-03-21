package todo

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/responses"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	ActivateTodoUseCase todo.ActivateTodoUseCase
	CompleteTodoUseCase todo.CompleteTodoUseCase
	CreateTodoUseCase   todo.CreateTodoUseCase
	DeleteTodoUseCase   todo.DeleteTodoUseCase
	GetTodoUseCase      todo.GetTodoUseCase
	ImportTodosUseCase  todo.ImportTodosUseCase
	ListTodosUseCase    todo.ListTodosUseCase
	RestartTodoUseCase  todo.RestartTodoUseCase
	StartTodoUseCase    todo.StartTodoUseCase
	UpdateTodoUseCase   todo.UpdateTodoUseCase
	logger              *log.Logger
}

func (h Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
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

	if payload.Recurring == false {
		if payload.TargetDate == nil {
			handlers.ThrowError(pkg.NilTargetDateError, http.StatusBadRequest, w)
			return
		}
	}

	if payload.Recurring == true {
		if payload.Recurrency == nil {
			handlers.ThrowError(pkg.NilRecurrencyError, http.StatusBadRequest, w)
			return
		}
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	t, err := h.CreateTodoUseCase.Execute(context.TODO(), userID, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(responses.CreateEntityResponse{ID: t.ID})
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
	b, err := json.Marshal(todos)
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
	// Get todo id & action (if present) from request URI
	path := strings.Split(r.RequestURI, handlers.TODO_STRING)[1]
	t := strings.Split(path, "/")[0]
	todoID, err := strconv.Atoi(t)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	if strings.Contains(r.RequestURI, handlers.START_STRING) {
		h.Start(w, r, uint(todoID))
	} else if strings.Contains(r.RequestURI, handlers.COMPLETE_STRING) {
		h.Complete(w, r, uint(todoID))
	} else if strings.Contains(r.RequestURI, handlers.RESTART_STRING) {
		h.Restart(w, r, uint(todoID))
	} else if strings.Contains(r.RequestURI, handlers.ACTIVATE_STRING) {
		h.Activate(w, r, uint(todoID))
	} else {
		switch r.Method {
		case http.MethodGet:
			h.Get(w, r, uint(todoID))
		case http.MethodDelete:
			h.Delete(w, r, uint(todoID))
		case http.MethodPut:
			h.Update(w, r, uint(todoID))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
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

	if payload.Recurring == false {
		if payload.TargetDate == nil {
			handlers.ThrowError(pkg.NilTargetDateError, http.StatusBadRequest, w)
			return
		}
	}

	if payload.Recurring == true {
		if payload.Recurrency == nil {
			handlers.ThrowError(pkg.NilRecurrencyError, http.StatusBadRequest, w)
			return
		}
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

func (h Handler) Start(w http.ResponseWriter, r *http.Request, id uint) {
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

	if err = h.StartTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Complete(w http.ResponseWriter, r *http.Request, id uint) {
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

	if err = h.CompleteTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Restart(w http.ResponseWriter, r *http.Request, id uint) {
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

	if err = h.RestartTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Activate(w http.ResponseWriter, r *http.Request, id uint) {
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

	if err = h.ActivateTodoUseCase.Execute(context.TODO(), payload, userID); err != nil {
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

func NewTodoHandler(
	activateTodoUseCase *todo.ActivateTodoUseCase,
	completeTodoUseCase *todo.CompleteTodoUseCase,
	createTodoUseCase *todo.CreateTodoUseCase,
	deleteTodoUseCase *todo.DeleteTodoUseCase,
	getTodoUseCase *todo.GetTodoUseCase,
	importTodosUseCase *todo.ImportTodosUseCase,
	listTodosUseCase *todo.ListTodosUseCase,
	restartTodoUseCase *todo.RestartTodoUseCase,
	startTodoUseCase *todo.StartTodoUseCase,
	updateTodoUseCase *todo.UpdateTodoUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		ActivateTodoUseCase: *activateTodoUseCase,
		CompleteTodoUseCase: *completeTodoUseCase,
		CreateTodoUseCase:   *createTodoUseCase,
		DeleteTodoUseCase:   *deleteTodoUseCase,
		GetTodoUseCase:      *getTodoUseCase,
		ImportTodosUseCase:  *importTodosUseCase,
		ListTodosUseCase:    *listTodosUseCase,
		StartTodoUseCase:    *startTodoUseCase,
		RestartTodoUseCase:  *restartTodoUseCase,
		UpdateTodoUseCase:   *updateTodoUseCase,
		logger:              logger,
	}
}
