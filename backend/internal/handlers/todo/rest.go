package todo

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

type TodoHandler struct {
	todoService ports.TodoServicer
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
		h.DeleteTodo(w, r, todoId)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	payload := ports.CreateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Create(context.TODO(), r, payload)
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

	err = h.todoService.Update(context.TODO(), r, payload)
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
	err = h.todoService.Complete(context.TODO(), r, payload)
	if err != nil {
		return
	}
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
	err = h.todoService.Activate(context.TODO(), r, payload)
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

	err = h.todoService.Start(context.TODO(), r, payload)
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

	err = h.todoService.Restart(context.TODO(), r, payload)
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

	todo, err := h.todoService.Get(context.TODO(), r, payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(todo)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func (h TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	if err != nil {
		return
	}
	payload := ports.ListTodosRequest{}
	payload.Category = categoryId
	todos, err := h.todoService.List(context.TODO(), r, payload)
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
}

func (h TodoHandler) ListRecurringTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListRecurring(context.TODO(), r)
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
}

func (h TodoHandler) ListCompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListCompleted(context.TODO(), r)
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
}

func (h TodoHandler) ListSuggestedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.ListSuggested(context.TODO(), r)
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
}

func (h TodoHandler) SuggestTodos(w http.ResponseWriter, r *http.Request) {
	err := h.todoService.Suggest(context.TODO(), r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	q := r.URL.Query()
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	if err != nil {
		return
	}
	payload := ports.DeleteTodoRequest{TodoId: int64(id), Category: categoryId}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.todoService.Delete(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h TodoHandler) Summary(w http.ResponseWriter, r *http.Request) {
	summary, err := h.todoService.Summary(context.TODO(), r)
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

func NewTodoHandler(ts ports.TodoServicer) TodoHandler {
	return TodoHandler{
		todoService: ts,
	}
}
