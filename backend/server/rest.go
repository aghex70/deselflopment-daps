package server

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/todo"
	"github.com/aghex70/daps/internal/handlers/user"
	"log"
	"net/http"
)

type RestServer struct {
	logger          *log.Logger
	cfg             config.RestConfig
	categoryHandler category.CategoryHandler
	toDoHandler     todo.TodoHandler
	userHandler     user.UserHandler
	categoryService ports.CategoryServicer
	toDoService     ports.TodoServicer
	userService     ports.UserServicer
}

func (s *RestServer) StartServer() error {
	// Categories
	http.HandleFunc("/categories", s.categoryHandler.ListCategories)
	http.HandleFunc("/category", s.categoryHandler.Category)

	// User
	http.HandleFunc("/login", s.userHandler.Login)
	http.HandleFunc("/logout", s.userHandler.Logout)
	http.HandleFunc("/register", s.userHandler.Register)
	http.HandleFunc("/refresh-token", s.userHandler.RefreshToken)
	http.HandleFunc("/user", s.userHandler.RemoveUser)

	// Todos
	http.HandleFunc("/todo", s.toDoHandler.Todo)
	http.HandleFunc("/todos", s.toDoHandler.ListTodos)

	// Stats
	//http.HandleFunc("/statistics", s.toDoHandler.ListTodos)

	address := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Starting server on address %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("Error starting server %+v", err.Error())
		return err
	}
	fmt.Println("Server started")
	return nil
}

func NewRestServer(cfg *config.RestConfig, ch category.CategoryHandler, tdh todo.TodoHandler, uh user.UserHandler, logger *log.Logger) *RestServer {
	return &RestServer{
		cfg:             *cfg,
		logger:          logger,
		categoryHandler: ch,
		toDoHandler:     tdh,
		userHandler:     uh,
	}
}
