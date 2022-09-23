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
	toDoService ports.TodoServicer
	logger      *log.Logger
}

func (h TodoHandler) Todo(w http.ResponseWriter, r *http.Request) {
	todoString := "todo/"
	if !strings.Contains(r.RequestURI, todoString) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	todoIdString := strings.Split(r.RequestURI, todoString)[1]
	todoId, err := strconv.Atoi(todoIdString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodDelete:
		h.DeleteTodo(w, r, todoId)
	case http.MethodGet:
		h.GetTodo(w, r, todoId)
	case http.MethodPut:
		h.UpdateTodo(w, r, todoId)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)

	payload := ports.CreateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Create(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	payload := ports.CompleteTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Complete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) StartTodo(w http.ResponseWriter, r *http.Request) {
	payload := ports.StartTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Start(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.DeleteTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Delete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.GetTodoRequest{TodoId: int64(id)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	todo, err := h.toDoService.Get(nil, r, payload)
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

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	payload := ports.UpdateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Update(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.toDoService.List(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.ListTodosResponse{Todos: todos})
	w.Write(b)
}

func NewTodoHandler(ts ports.TodoServicer, logger *log.Logger) TodoHandler {
	return TodoHandler{
		toDoService: ts,
		logger:      logger,
	}
}
