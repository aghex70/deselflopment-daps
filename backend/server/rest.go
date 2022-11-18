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
	todoService     ports.TodoServicer
	userService     ports.UserServicer
	rootHandler     root.RootHandler
}

var hmacSampleSecret = []byte("random")

func JWTAuthMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT, OPTIONS")
		//w.Header().Add("Access-Control-Allow-Credentials", "true")
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
	fmt.Printf("%+v", payload)
	value := reflect.ValueOf(payload)
	bodyToken := value.FieldByName("AccessToken").String()
	return bodyToken
}

func RetrieveJWTClaims(r *http.Request, payload interface{}) (float64, error) {
	fmt.Printf("\n\npayload -------------> %+v", payload)
	fmt.Printf("\n\nrequest -------------> %+v", r)
	var tokenString string
	fmt.Println("\n\nAUTHORIZATION -------------------> ", r.Header["Authorization"])
	if r.Header["Authorization"] != nil {
		tokenString = RetrieveHeaderJWT(r)
	} else {
		tokenString = RetrieveBodyJWT(payload)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	fmt.Printf("\n\nCLAIMS -------------------> %+v", claims)
	userId := claims["user_id"].(float64)
	return userId, nil
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("Access-Control-Allow-Origin", "*")
		//w.Header().Add("Access-Control-Allow-Credentials", "true")
		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//
		//if r.Method == "OPTIONS" {
		//	http.Error(w, "No Content", http.StatusNoContent)
		//	return
		//}

		next(w, r)
	}
}

func (s *RestServer) StartServer() error {
	// User
	http.HandleFunc("/register", s.userHandler.Register)
	http.HandleFunc("/login", s.userHandler.Login)
	http.HandleFunc("/refresh-token", JWTAuthMiddleware(s.userHandler.RefreshToken))
	http.HandleFunc("/user", JWTAuthMiddleware(s.userHandler.RemoveUser))
	//http.HandleFunc("/recover-password", JWTAuthMiddleware(s.userHandler.RemoveUser))

	// Categories
	http.HandleFunc("/categories", s.categoryHandler.ListCategories)
	http.HandleFunc("/category", s.categoryHandler.CreateCategory)

	// Todos
	http.HandleFunc("/todo", JWTAuthMiddleware(s.toDoHandler.CreateTodo))
	http.HandleFunc("/todos", JWTAuthMiddleware(s.toDoHandler.ListTodos))
	http.HandleFunc("/recurring-todos", JWTAuthMiddleware(s.toDoHandler.ListRecurringTodos))
	http.HandleFunc("/completed-todos", JWTAuthMiddleware(s.toDoHandler.ListCompletedTodos))
	http.HandleFunc("/suggested-todos", JWTAuthMiddleware(s.toDoHandler.ListCompletedTodos))
	http.HandleFunc("/summary", JWTAuthMiddleware(s.toDoHandler.Summary))

	// Stats
	//http.HandleFunc("/statistics", JWTAuthMiddleware(s.toDoHandler.Todo))

	// CAREFUL!!!!
	// Root (not included out of the box damn!)
	http.HandleFunc("/", JWTAuthMiddleware(s.rootHandler.Root))

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

func NewRestServer(cfg *config.RestConfig, ch category.CategoryHandler, tdh todo.TodoHandler, uh user.UserHandler, rh root.RootHandler, logger *log.Logger) *RestServer {
	return &RestServer{
		cfg:             *cfg,
		logger:          logger,
		categoryHandler: ch,
		toDoHandler:     tdh,
		userHandler:     uh,
		rootHandler:     rh,
	}
}
