package server

import (
	"fmt"
	"github.com/aghex70/daps/config"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"github.com/aghex70/daps/internal/handlers/category"
	"github.com/aghex70/daps/internal/handlers/root"
	"github.com/aghex70/daps/internal/handlers/todo"
	"github.com/aghex70/daps/internal/handlers/user"
	"github.com/aghex70/daps/internal/handlers/userconfig"
	"github.com/aghex70/daps/pkg"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type RestServer struct {
	logger            *log.Logger
	cfg               config.RestConfig
	categoryHandler   category.CategoryHandler
	toDoHandler       todo.TodoHandler
	userHandler       user.UserHandler
	categoryService   ports.CategoryServicer
	todoService       ports.TodoServicer
	userService       ports.UserServicer
	rootHandler       root.RootHandler
	userConfigHandler userconfig.UserConfigHandler
}

func JWTAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", pkg.GetOrigin())
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT, OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

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
			return pkg.HmacSampleSecret, nil
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
		return pkg.HmacSampleSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)
	return userId, nil
}

func (s *RestServer) StartServer() error {
	// User
	http.HandleFunc("/api/register", s.userHandler.Register)
	http.HandleFunc("/api/login", s.userHandler.Login)
	http.HandleFunc("/api/refresh-token", JWTAuthMiddleware(s.userHandler.RefreshToken))
	http.HandleFunc("/api/reset-link", s.userHandler.ResetLink)
	http.HandleFunc("/api/reset-password", s.userHandler.ResetPassword)
	http.HandleFunc("/api/users", JWTAuthMiddleware(s.userHandler.ListUsers))
	http.HandleFunc("/api/user/admin", JWTAuthMiddleware(s.userHandler.CheckAdmin))
	http.HandleFunc("/api/user/provision", JWTAuthMiddleware(s.userHandler.ProvisionDemoUser))
	http.HandleFunc("/api/user/activate", s.userHandler.ActivateUser)

	// Categories
	http.HandleFunc("/api/categories", JWTAuthMiddleware(s.categoryHandler.ListCategories))
	http.HandleFunc("/api/category", JWTAuthMiddleware(s.categoryHandler.CreateCategory))

	// Todos
	http.HandleFunc("/api/todo", JWTAuthMiddleware(s.toDoHandler.CreateTodo))
	http.HandleFunc("/api/todos", JWTAuthMiddleware(s.toDoHandler.ListTodos))
	http.HandleFunc("/api/recurring-todos", JWTAuthMiddleware(s.toDoHandler.ListRecurringTodos))
	http.HandleFunc("/api/completed-todos", JWTAuthMiddleware(s.toDoHandler.ListCompletedTodos))
	http.HandleFunc("/api/suggest", JWTAuthMiddleware(s.toDoHandler.SuggestTodos))
	http.HandleFunc("/api/suggested-todos", JWTAuthMiddleware(s.toDoHandler.ListSuggestedTodos))
	http.HandleFunc("/api/summary", JWTAuthMiddleware(s.toDoHandler.Summary))
	http.HandleFunc("/api/remind", s.toDoHandler.Remind)

	// UserConfiguration
	http.HandleFunc("/api/user-configuration/", JWTAuthMiddleware(s.userConfigHandler.HandleUserConfig))

	// CSV Import
	http.HandleFunc("/api/import", JWTAuthMiddleware(s.userHandler.ImportCSV))

	// CAREFUL!!!!
	// Root (not included out of the box damn!)
	http.HandleFunc("/api/", JWTAuthMiddleware(s.rootHandler.Root))

	address := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	fmt.Printf("Starting server on address %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("Error starting HTTP server %+v", err.Error())
		return err
	}

	fmt.Println("Server started")
	return nil
}

func NewRestServer(cfg *config.RestConfig, ch category.CategoryHandler, tdh todo.TodoHandler, uh user.UserHandler, rh root.RootHandler, uch userconfig.UserConfigHandler, logger *log.Logger) *RestServer {
	return &RestServer{
		cfg:               *cfg,
		logger:            logger,
		categoryHandler:   ch,
		toDoHandler:       tdh,
		userHandler:       uh,
		rootHandler:       rh,
		userConfigHandler: uch,
	}
}
