package user

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/user"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/server"
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

func (s Service) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.userRepository.GetByEmail(ctx, email)
}

func (s Service) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return domain.User{}, err
	}

	return nu, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	_, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Activate(ctx context.Context, activationCode string) error {
	return s.userRepository.Activate(ctx, activationCode)
}

func (s Service) Update(ctx context.Context, fields map[string]interface{}) error {
	//err := s.userRepository.Update(ctx, activationCode)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s Service) CheckAdmin(ctx context.Context, r *http.Request) (uint, error) {
	userID, err := server.RetrieveJWTClaims(r, nil)
	if err != nil {
		return 0, pkg.InvalidCredentialsError
	}
	u, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		return 0, pkg.InvalidCredentialsError
	}

	if !u.Admin {
		return 0, pkg.UnauthorizedError
	}

	return userID, nil
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
	//	err, _ = s.repository.CreateTodo(ctx, domain.Todo{
	//		Name:       name,
	//		Link:       link,
	//		CategoryID: uint(categoryID),
	//		Priority:   domain.Priority(3),
	//	})
	//	if err != nil {
	//		return err
	//	}
	//}

	return nil
}

func (s Service) List(ctx context.Context, r *http.Request) ([]domain.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//users, err := s.repository.GetUsers(ctx)
	//if err != nil {
	//	return []domain.User{}, err
	//}
	//
	//return users, nil
	return []domain.User{}, nil
}

func (s Service) ProvisionDemoUser(ctx context.Context, r *http.Request, req requests.ProvisionDemoUserRequest) error {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//cipheredPassword := s.EncryptPassword(ctx, req.Password)
	//u := domain.User{
	//	Name:              pkg2.DemoUserName,
	//	Email:             req.Email,
	//	Password:          cipheredPassword,
	//	Active:            true,
	//	ResetPasswordCode: pkg2.GenerateUUID(),
	//}
	//
	//nu, err := s.repository.CreateUser(ctx, u)
	//if err != nil {
	//	return err
	//}
	//
	////nuc := domain.UserConfig{
	////	UserID:      nu.ID,
	////	AutoSuggest: false,
	////	Language:    "en",
	////}
	//
	////err = s.userConfigurationRepository.Create(ctx, nuc)
	////err = s.userConfigurationRepository.Create(ctx, nuc)
	////if err != nil {
	////	return err
	////}
	//
	//demoCategory := domain.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Home tasks",
	//	Custom: true,
	//	Name:   "Home",
	//	//Users:       []domain.User{u},
	//}
	//
	//c, err := s.repository.CreateCategory(ctx, demoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//anotherDemoCategory := domain.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Work stuff",
	//	Custom: true,
	//	Name:   "Work",
	//	//Users:       []domain.User{u},
	//}
	//
	//ac, err := s.repository.CreateCategory(ctx, anotherDemoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//yetAnotherDemoCategory := domain.Category{
	//	OwnerID: nu.ID,
	//	//Description: "Purchase list",
	//	Custom: true,
	//	Name:   "Purchases",
	//	//Users:       []domain.User{u},
	//}
	//
	//yac, err := s.repository.CreateCategory(ctx, yetAnotherDemoCategory)
	//if err != nil {
	//	return err
	//}
	//
	//todos := pkg2.GenerateDemoTodos(int(c.ID), int(ac.ID), int(yac.ID), req.Language)
	//
	//for _, t := range todos {
	//	_, _ = s.repository.CreateTodo(ctx, t)
	//	//if err != nil {
	//	//	return err
	//	//}
	//}

	return nil
}

func (s Service) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	//userID, err := server.RetrieveJWTClaims(r, nil)
	//if err != nil {
	//	return "", errors.New("invalid token")
	//}
	//u, err := s.repository.GetUser(ctx, uint(int(userID)))
	//if err != nil {
	//	return "", errors.New("invalid token")
	//}
	//
	//newClaims := MyCustomClaims{
	//	UserID: u.ID,
	//	Admin:  u.Admin,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		Subject:   u.Email,
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
	//	},
	//}
	//
	//mySigningKey := pkg2.HmacSampleSecret
	//newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	//ss, err := newToken.SignedString(mySigningKey)
	//if err != nil {
	//	return "", err
	//}
	//
	//return ss, nil
	return "", nil
}

func (s Service) ResetPassword(ctx context.Context, r requests.ResetPasswordRequest) error {
	//match := s.PasswordMatchesRepeatPassword(ctx, r.Password, r.RepeatPassword)
	//if !match {
	//	return errors.New("passwords do not match")
	//}
	//
	//encryptedPassword := s.EncryptPassword(ctx, r.Password)
	//err := s.repository.ResetPassword(ctx, encryptedPassword, r.ResetPasswordCode)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request, req requests.GetUserRequest) (domain.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return domain.User{}, err
	//}
	//
	//u, err := s.repository.GetUser(ctx, uint(int(req.UserID)))
	//if err != nil {
	//	return domain.User{}, err
	//}
	//
	//return u, nil
	return domain.User{}, nil
}

func (s Service) SendResetLink(ctx context.Context, r requests.ResetLinkRequest) error {
	//u, err := s.repository.CreateResetLink(ctx, r.Email)
	//if err != nil {
	//	return err
	//}
	//
	//e := domain.Email{
	//	From:      pkg.FromEmail,
	//	To:        u.Email,
	//	Recipient: u.Name,
	//	Subject:   "📣 DAPS - Password reset request 📣",
	//	Body:      "In order to reset your password, please follow this link: " + pkg.ResetPasswordLink + u.ResetPasswordCode,
	//	User:      u.ID,
	//}
	//
	//err = pkg.SendEmail(e)
	//if err != nil {
	//	fmt.Printf("Error sending email: %+v", err)
	//	e.Error = err.Error()
	//	e.Sent = false
	//	_, errz := s.repository.CreateEmail(ctx, e)
	//	if errz != nil {
	//		fmt.Printf("Error saving email: %+v", errz)
	//		return errz
	//	}
	//	return err
	//}
	//
	//e.Sent = true
	//_, err = s.repository.CreateEmail(ctx, e)
	//if err != nil {
	//	return err
	//}
	return nil
}

func NewUserService(ur user.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		userRepository: ur,
	}
}
