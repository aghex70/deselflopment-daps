package server

import (
	"fmt"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/handlers/category"
	"github.com/aghex70/daps/internal/ports/handlers/todo"
	"github.com/aghex70/daps/internal/ports/handlers/user"
	"log"
	"net/http"
	"time"

	"github.com/aghex70/daps/config"
)

type RestServer struct {
	logger          *log.Logger
	cfg             config.RestConfig
	categoryHandler category.Handler
	toDoHandler     todo.Handler
	userHandler     user.Handler
}

func (s *RestServer) StartServer() error {
	// User
	http.HandleFunc("/api/users", handlers.JWTAuthMiddleware(s.userHandler.List))
	http.HandleFunc("/api/users/provision", handlers.JWTAuthMiddleware(s.userHandler.ProvisionDemoUser))
	http.HandleFunc("/api/register", s.userHandler.Register)
	http.HandleFunc("/api/login", s.userHandler.Login)
	http.HandleFunc("/api/refresh-token", handlers.JWTAuthMiddleware(s.userHandler.RefreshToken))
	http.HandleFunc("/api/reset-link", handlers.JWTAuthMiddleware(s.userHandler.ResetLink))
	http.HandleFunc("/api/reset-password", handlers.JWTAuthMiddleware(s.userHandler.ResetPassword))
	http.HandleFunc("/api/activate", handlers.JWTAuthMiddleware(s.userHandler.Activate))
	http.HandleFunc("/api/users/", handlers.JWTAuthMiddleware(s.userHandler.HandleUser))

	// Categories
	http.HandleFunc("/api/categories", handlers.JWTAuthMiddleware(s.categoryHandler.HandleCategories))
	http.HandleFunc("/api/categories/", handlers.JWTAuthMiddleware(s.categoryHandler.HandleCategory))

	//// Todos
	http.HandleFunc("/api/todos", handlers.JWTAuthMiddleware(s.toDoHandler.HandleTodos))
	http.HandleFunc("/api/todos/", handlers.JWTAuthMiddleware(s.toDoHandler.HandleTodo))
	http.HandleFunc("/api/todos/import", handlers.JWTAuthMiddleware(s.toDoHandler.Import))
	//http.HandleFunc("/api/suggest", JWTAuthMiddleware(s.toDoHandler.SuggestTodos))
	////http.HandleFunc("/api/summary", JWTAuthMiddleware(s.toDoHandler.Summary))

	address := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	log.Printf("Starting server on address %s", address)
	server := &http.Server{
		Addr:              address,
		ReadHeaderTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error starting HTTP server %+v", err.Error())
		return err
	}

	log.Println("Server started")
	return nil
}

func NewRestServer(cfg *config.RestConfig, ch category.Handler, tdh todo.Handler, uh user.Handler, logger *log.Logger) *RestServer {
	return &RestServer{
		cfg:             *cfg,
		logger:          logger,
		categoryHandler: ch,
		toDoHandler:     tdh,
		userHandler:     uh,
	}
}
