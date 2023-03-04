package todo

import (
	"encoding/json"
	"errors"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type TodoHandler struct {
	todoService ports.TodoServicer
	logger      *log.Logger
}

func (h TodoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.TODO_STRING)[1]
	if startString := "/start"; strings.Contains(path, startString) {
		todoIdString := strings.Split(path, startString)[0]
		todoId, err := strconv.Atoi(todoIdString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.StartTodo(w, r, todoId, 0)
		return
	}

	if restartString := "/restart"; strings.Contains(path, restartString) {
		todoIdString := strings.Split(path, restartString)[0]
		todoId, err := strconv.Atoi(todoIdString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.RestartTodo(w, r, todoId, 0)
		return
	}

	if completeString := "/complete"; strings.Contains(path, completeString) {
		todoIdString := strings.Split(path, completeString)[0]
		todoId, err := strconv.Atoi(todoIdString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.CompleteTodo(w, r, todoId, 0)
		return
	}

	if activateString := "/activate"; strings.Contains(path, activateString) {
		todoIdString := strings.Split(path, activateString)[0]
		todoId, err := strconv.Atoi(todoIdString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.ActivateTodo(w, r, todoId, 0)
		return
	}

	queryParams := strings.Split(path, "?")
	todoId, err := strconv.Atoi(queryParams[0])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(queryParams) == 1 {
		if r.Method == http.MethodPut {
			h.UpdateTodo(w, r, todoId)
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TodoId and CategoryId
	cId := strings.Split(queryParams[1], "=")[1]
	categoryId, err := strconv.Atoi(cId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		h.GetTodo(w, r, todoId, categoryId)
		return
	}

	if r.Method == http.MethodDelete {
		h.DeleteTodo(w, r, todoId, categoryId)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	payload := ports.CreateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Create(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.UpdateTodoRequest{TodoId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Update(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.CompleteTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	err = h.todoService.Complete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) ActivateTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.ActivateTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	err = h.todoService.Activate(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) StartTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.StartTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Start(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) RestartTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.StartTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Restart(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.GetTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	todo, err := h.todoService.Get(nil, r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(todo)
	w.Write(b)
}

func (h TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	payload := ports.ListTodosRequest{}
	payload.Category = categoryId
	todos, err := h.todoService.List(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) ListRecurringTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListRecurring(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) ListCompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListCompleted(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) ListSuggestedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListSuggested(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) SuggestTodos(w http.ResponseWriter, r *http.Request) {
	err := h.todoService.Suggest(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	q := r.URL.Query()
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	payload := ports.DeleteTodoRequest{TodoId: int64(id), Category: categoryId}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Delete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h TodoHandler) Summary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.todoService.Summary(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	// Return a list of summaries
	b, err := json.Marshal(summary)
	w.Write(b)
}

func (h TodoHandler) Remind(w http.ResponseWriter, r *http.Request) {
	err := h.todoService.Remind(nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewTodoHandler(ts ports.TodoServicer, logger *log.Logger) TodoHandler {
	return TodoHandler{
		todoService: ts,
		logger:      logger,
	}
}
