package todo

import (
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/todo"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	//todoService todo.Servicer
}

func (h Handler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.RequestURI, handlers.TODO_STRING)[1]
	if startString := "/start"; strings.Contains(path, startString) {
		todoIDString := strings.Split(path, startString)[0]
		todoID, err := strconv.Atoi(todoIDString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.StartTodo(w, r, uint(todoID), 0)
		return
	}

	if restartString := "/restart"; strings.Contains(path, restartString) {
		todoIDString := strings.Split(path, restartString)[0]
		todoID, err := strconv.Atoi(todoIDString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.RestartTodo(w, r, uint(todoID), 0)
		return
	}

	if completeString := "/complete"; strings.Contains(path, completeString) {
		todoIDString := strings.Split(path, completeString)[0]
		todoID, err := strconv.Atoi(todoIDString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.CompleteTodo(w, r, uint(todoID), 0)
		return
	}

	if activateString := "/activate"; strings.Contains(path, activateString) {
		todoIDString := strings.Split(path, activateString)[0]
		todoID, err := strconv.Atoi(todoIDString)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// TODO PENDING
		h.ActivateTodo(w, r, uint(todoID), 0)
		return
	}

	queryParams := strings.Split(path, "?")
	todoID, err := strconv.Atoi(queryParams[0])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if len(queryParams) == 1 {
		if r.Method == http.MethodPut {
			h.UpdateTodo(w, r, uint(todoID))
			return
		}

		w.WriteHeader(http.StatusNotFound)
		return
	}

	// TodoID and CategoryID
	cID := strings.Split(queryParams[1], "=")[1]
	//categoryID, err := strconv.Atoi(cID)
	_, err = strconv.Atoi(cID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		//h.GetTodo(w, r, uint(todoID), uint(categoryID))
		return
	}

	if r.Method == http.MethodDelete {
		h.DeleteTodo(w, r, uint(todoID))
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	payload := requests.CreateTodoRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.todoService.Create(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) UpdateTodo(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateTodoRequest{TodoID: uint(int64(id))}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.todoService.Update(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	return
}

func (h Handler) CompleteTodo(w http.ResponseWriter, r *http.Request, id, categoryID uint) {
	payload := requests.CompleteTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	//err = h.todoService.Complete(context.TODO(), r, payload)
	//if err != nil {
	//	return
	//}
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	return
}

func (h Handler) ActivateTodo(w http.ResponseWriter, r *http.Request, id, categoryID uint) {
	payload := requests.ActivateTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	//err = h.todoService.Activate(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	return
}

func (h Handler) StartTodo(w http.ResponseWriter, r *http.Request, id, categoryID uint) {
	payload := requests.StartTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.todoService.Start(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	return
}

func (h Handler) RestartTodo(w http.ResponseWriter, r *http.Request, id, categoryID uint) {
	payload := requests.StartTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.todoService.Restart(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	return
}

//func (h Handler) GetTodo(w http.ResponseWriter, r *http.Request, id, categoryID uint) {
//	payload := requests.GetTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
//	err := handlers.ValidateRequest(r, &payload)
//	if err != nil {
//		handlers.ThrowError(err, http.StatusBadRequest, w)
//		return
//	}
//
//	todo, err := h.todoService.Get(context.TODO(), r, payload)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		handlers.ThrowError(err, http.StatusBadRequest, w)
//		return
//	}
//	b, err := json.Marshal(todo)
//	if err != nil {
//		return
//	}
//	_, err = w.Write(b)
//	if err != nil {
//		return
//	}
//}

func (h Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	categoryID, err := strconv.Atoi(q.Get("category_id"))
	if err != nil {
		return
	}
	payload := requests.ListTodosRequest{}
	payload.Category = uint(categoryID)
	//todos, err := h.todoService.List(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//
	//b, err := json.Marshal(todos)
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
	return
}

func (h Handler) ListRecurringTodos(w http.ResponseWriter, r *http.Request) {
	//todos, err := h.todoService.ListRecurring(context.TODO(), r)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//
	//b, err := json.Marshal(todos)
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
	return
}

func (h Handler) ListCompletedTodos(w http.ResponseWriter, r *http.Request) {
	//todos, err := h.todoService.ListCompleted(context.TODO(), r)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	//
	//b, err := json.Marshal(todos)
	//if err != nil {
	//	return
	//}
	//_, err = w.Write(b)
	//if err != nil {
	//	return
	//}
	return
}

//func (h Handler) ListSuggestedTodos(w http.ResponseWriter, r *http.Request) {
//	todos, err := h.todoService.ListSuggested(context.TODO(), r)
//	if err != nil {
//		handlers.ThrowError(err, http.StatusBadRequest, w)
//		return
//	}
//
//	b, err := json.Marshal(todos)
//	if err != nil {
//		return
//	}
//	_, err = w.Write(b)
//	if err != nil {
//		return
//	}
//}

func (h Handler) SuggestTodos(w http.ResponseWriter, r *http.Request) {
	//err := h.todoService.Suggest(context.TODO(), r)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteTodo(w http.ResponseWriter, r *http.Request, id uint) {
	q := r.URL.Query()
	categoryID, err := strconv.Atoi(q.Get("category_id"))
	if err != nil {
		return
	}
	payload := requests.DeleteTodoRequest{TodoID: uint(int64(id)), Category: uint(categoryID)}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	//err = h.todoService.Delete(context.TODO(), r, payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}
	w.WriteHeader(http.StatusNoContent)
}

//func (h Handler) Summary(w http.ResponseWriter, r *http.Request) {
//	summary, err := h.todoService.Summary(context.TODO(), r)
//	if err != nil {
//		handlers.ThrowError(err, http.StatusBadRequest, w)
//		return
//	}
//
//	b, err := json.Marshal(summary)
//	if err != nil {
//		return
//	}
//	_, err = w.Write(b)
//	if err != nil {
//		return
//	}
//}

//func NewTodoHandler(ts todo.Servicer) Handler {
//	return Handler{
//		todoService: ts,
//	}
//}
