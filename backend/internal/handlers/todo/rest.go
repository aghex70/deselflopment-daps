package todo

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

type TodoHandler struct {
	toDoService ports.TodoServicer
	logger      *log.Logger
}

func (h TodoHandler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.TODO_STRING)[1]
	fmt.Println("path", path)

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

	if completeString := "/activate"; strings.Contains(path, completeString) {
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
		h.ActivateTodo(w, r, todoId, 0)
		return
	}

	queryParams := strings.Split(path, "?")
	todoId, err := strconv.Atoi(queryParams[0])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Println("queryParams -------------> ", queryParams)
	if len(queryParams) == 1 {
		if r.Method == http.MethodPut {
			fmt.Println("PUT")
			fmt.Printf("\n\nrequest: %+v", r)
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
		fmt.Println("GET")
		fmt.Printf("\n\nrequest: %+v", r)
		h.GetTodo(w, r, todoId, categoryId)
		return
	}

	if r.Method == http.MethodDelete {
		fmt.Println("DELETE")
		h.DeleteTodo(w, r, todoId, categoryId)
		return
	}

	fmt.Println("Hola5")
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}

func (h TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("asdadadasd")
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
	w.WriteHeader(http.StatusCreated)
}

func (h TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Println("request: ", r)
	fmt.Println("request: ", r.Body)
	payload := ports.UpdateTodoRequest{TodoId: int64(id)}
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

func (h TodoHandler) CompleteTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.CompleteTodoRequest{TodoId: int64(id), Category: categoryId}
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

func (h TodoHandler) ActivateTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	payload := ports.ActivateTodoRequest{TodoId: int64(id), Category: categoryId}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	err = h.toDoService.Activate(nil, r, payload)
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

	err = h.toDoService.Start(nil, r, payload)
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

func (h TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\nrequest ------>: %+v", r)
	q := r.URL.Query()
	bod := r.Body
	fmt.Printf("\n\nqparams ------>: %+v", q)
	fmt.Printf("\n\nbody ------>: %+v", bod)
	payload := ports.ListTodosRequest{}
	//err := handlers.ValidateRequest(r, &payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}

	fmt.Println("alberto0")
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	fmt.Println("alberto1")
	payload.Category = categoryId
	fmt.Println("alberto")
	todos, err := h.toDoService.List(nil, r, payload)
	if err != nil {
		fmt.Println("alberto2")
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	fmt.Println("alberto3")
	w.Write(b)
}

func (h TodoHandler) ListRecurringTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.toDoService.ListRecurring(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) ListCompletedTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.toDoService.ListCompleted(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(todos)
	w.Write(b)
}

func (h TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id, categoryId int) {
	fmt.Printf("\n\nrequest ------>: %+v", r)
	q := r.URL.Query()
	bod := r.Body
	fmt.Printf("\n\nqparams ------>: %+v", q)
	fmt.Printf("\n\nbody ------>: %+v", bod)
	//payloadz := ports.ListTodosRequest{}
	//err := handlers.ValidateRequest(r, &payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}

	categoryId, err := strconv.Atoi(q.Get("category_id"))
	//payload.Category = categoryId
	payload := ports.DeleteTodoRequest{TodoId: int64(id), Category: categoryId}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.toDoService.Delete(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h TodoHandler) Summary(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//bodyBytes, _ := ioutil.ReadAll(r.Body)
	//bodyString := string(bodyBytes)
	//fmt.Println(bodyString)
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Credentials", "true")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	summary, err := h.toDoService.Summary(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	// Return a list of summaries
	b, err := json.Marshal(summary)
	//b, err := json.Marshal(handlers.SummaryResponse{Summary: summary})
	w.Write(b)
	//w.Write(summary)
}

func NewTodoHandler(ts ports.TodoServicer, logger *log.Logger) TodoHandler {
	return TodoHandler{
		toDoService: ts,
		logger:      logger,
	}
}
