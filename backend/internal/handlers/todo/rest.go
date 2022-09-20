package todo

import (
	"encoding/json"
	"fmt"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type TodoHandler struct {
	toDoService ports.TodoServicer
	logger      *log.Logger
}

func (h TodoHandler) Todo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodDelete:
		h.DeleteTodo(w, r)
	case http.MethodGet:
		h.GetTodo(w, r)
	case http.MethodPost:
		h.CreateTodo(w, r)
	case http.MethodPut:
		h.CompleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
		handlers.ThrowError(err, w)
		return
	}

	err = h.toDoService.Create(nil, payload)
	if err != nil {
		handlers.ThrowError(err, w)
		return
	}
}

func (h TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	payload := ports.ListTodosRequest{TodoId: 5}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, w)
		return
	}

	todos, err := h.toDoService.List(nil, payload)
	if err != nil {
		handlers.ThrowError(err, w)
		return
	}

	var response handlers.ListTodosResponse
	fmt.Println("response", response, "\ntodos", todos)
	b, err := json.Marshal(handlers.ListTodosResponse{Todos: todos})
	w.Write(b)
	//fmt.Fprint(w, string(b))
}

func NewTodoHandler(ts ports.TodoServicer, logger *log.Logger) TodoHandler {
	return TodoHandler{
		toDoService: ts,
		logger:      logger,
	}
}
