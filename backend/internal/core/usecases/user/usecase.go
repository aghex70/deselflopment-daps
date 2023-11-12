package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/repositories/user"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"mime/multipart"
	"net/http"
)

type Service struct {
	logger         *log.Logger
	userRepository user.Repository
}

type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func (s Service) ImportCSV(ctx context.Context, r *http.Request, f multipart.File) error {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//// Create a buffer to read the file line by line
	//buf := bufio.NewReader(f)
	//
	//// Parse the CSV file
	//rr := csv.NewReader(buf)
	//
	//// Read and discard the first line
	//_, err = rr.Read()
	//if err != nil {
	//	return err
	//}
	//
	//// Iterate over the lines of the CSV file
	//for {
	//	// Read the next line
	//	record, err := rr.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		return err
	//	}
	//
	//	name := record[0]
	//	link := record[1]
	//	categoryID, _ := strconv.Atoi(record[2])
	//
	//	err, _ = s.repository.CreateTodo(ctx, domain2.Todo{
	//		Name:       name,
	//		Link:       link,
	//		CategoryID: uint(categoryID),
	//		Priority:   domain2.Priority(3),
	//	})
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func NewUserService(ur user.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		userRepository: ur,
	}
}
