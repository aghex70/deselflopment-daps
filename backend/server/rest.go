package server

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/todo"
	"github.com/aghex70/daps/internal/handlers/user"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
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

var hmacSampleSecret = []byte("random")

func JWTAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorizationHeader := r.Header["Authorization"][0]
		headerToken := strings.Split(authorizationHeader, " ")[1]
		token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})

		if err != nil {
			handlers.ThrowError(err, http.StatusUnauthorized, w)
			return
		}
		if !token.Valid {
			handlers.ThrowError(err, http.StatusUnauthorized, w)
			return
		}
		f(w, r)
	}
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
	http.HandleFunc("/todo", JWTAuthMiddleware(s.toDoHandler.Todo))
	http.HandleFunc("/todos", JWTAuthMiddleware(s.toDoHandler.ListTodos))

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
