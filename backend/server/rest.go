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
	"reflect"
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
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

func RetrieveHeaderJWT(r *http.Request) string {
	authorizationHeader := r.Header["Authorization"][0]
	headerToken := strings.Split(authorizationHeader, " ")[1]
	return headerToken
}

func RetrieveBodyJWT(payload interface{}) string {
	value := reflect.ValueOf(payload)
	bodyToken := value.FieldByName("AccessToken").String()
	return bodyToken
}

func RetrieveJWTClaims(r *http.Request, payload interface{}) (float64, error) {
	var tokenString string
	if r.Header["Authorization"] != nil {
		tokenString = RetrieveHeaderJWT(r)
	} else {
		tokenString = RetrieveBodyJWT(payload)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	return userId, nil
}

func (s *RestServer) StartServer() error {
	// Categories
	http.HandleFunc("/categories", s.categoryHandler.ListCategories)
	http.HandleFunc("/category", s.categoryHandler.Category)

	// User
	http.HandleFunc("/login", s.userHandler.Login)
	http.HandleFunc("/logout", JWTAuthMiddleware(s.userHandler.Logout))
	http.HandleFunc("/register", s.userHandler.Register)
	http.HandleFunc("/refresh-token", JWTAuthMiddleware(s.userHandler.RefreshToken))
	http.HandleFunc("/user", JWTAuthMiddleware(s.userHandler.RemoveUser))

	// Todos
	http.HandleFunc("/todo", JWTAuthMiddleware(s.toDoHandler.CreateTodo))
	http.HandleFunc("/todo/complete", JWTAuthMiddleware(s.toDoHandler.CompleteTodo))
	http.HandleFunc("/todo/start", JWTAuthMiddleware(s.toDoHandler.StartTodo))
	http.HandleFunc("/todos", JWTAuthMiddleware(s.toDoHandler.ListTodos))

	// Stats
	http.HandleFunc("/statistics", JWTAuthMiddleware(s.toDoHandler.Todo))

	// CAREFUL!!!!
	// Root (not included out of the box damn!)
	http.HandleFunc("/", JWTAuthMiddleware(s.toDoHandler.Todo))

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
