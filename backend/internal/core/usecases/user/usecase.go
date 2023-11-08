package user

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/user"
	"github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"mime/multipart"
	"net/http"
)

type Service struct {
	logger     *log.Logger
	repository user.Repository
}

type MyCustomClaims struct {
	UserID uint `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func (s Service) Register(ctx context.Context, r requests.CreateUserRequest) error {
	//preexistent := s.CheckExistentUser(ctx, r.Email)
	//if preexistent {
	//	return errors.New("user already registered")
	//}
	//cipheredPassword := s.EncryptPassword(ctx, r.Password)
	//
	//categories, err := s.repository.GetCategories(ctx, &pkg2.BaseCategoriesIdsFilter)
	//if err != nil {
	//	return err
	//}
	//u := domain2.User{
	//	Name:              r.Name,
	//	Email:             r.Email,
	//	Password:          cipheredPassword,
	//	Categories:        &categories,
	//	Active:            false,
	//	ActivationCode:    pkg2.GenerateUUID(),
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
	////if err != nil {
	////	return err
	////}
	//
	//e := domain2.Email{
	//	From:      pkg2.FromEmail,
	//	To:        r.Email,
	//	Recipient: r.Name,
	//	Subject:   "📣 DAPS - Activate your account 📣",
	//	Body:      "In order to complete your registration, please click on the following link: " + pkg2.ActivationCodeLink + nu.ActivationCode,
	//	UserID:    nu.ID,
	//}
	//
	//err = pkg2.SendEmail(e)
	//if err != nil {
	//	fmt.Printf("Error sending email: %+v", err)
	//	//e.Error = err.Error()
	//	e.Sent = false
	//
	//	err = s.repository.DeleteUser(ctx, nu.ID)
	//	if err != nil {
	//		return err
	//	}
	//
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

func (s Service) Login(ctx context.Context, r requests.LoginUserRequest) (string, int, error) {
	//u, err := s.repository.GetUserByEmail(ctx, r.Email)
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//decryptedPassword, err := s.DecryptPassword(ctx, u.Password)
	//
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//match := s.PasswordsMatch(ctx, decryptedPassword, r.Password)
	//if !match {
	//	return "", 0, errors.New("invalid credentials")
	//}
	//
	//claims := MyCustomClaims{
	//	UserID: u.ID,
	//	Admin:  u.Admin,
	//	RegisteredClaims: jwt.RegisteredClaims{
	//		Subject:   r.Email,
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
	//	},
	//}
	//
	//mySigningKey := pkg2.HmacSampleSecret
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//ss, err := token.SignedString(mySigningKey)
	//if err != nil {
	//	return "", 0, err
	//}
	//
	//return ss, int(u.ID), nil
	return "", 0, nil
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

func (s Service) CheckAdmin(ctx context.Context, r *http.Request) (int, error) {
	//userID, err := server.RetrieveJWTClaims(r, nil)
	//if err != nil {
	//	return 0, errors.New("invalid token")
	//}
	//u, err := s.repository.GetUser(ctx, uint(int(userID)))
	//if err != nil {
	//	return 0, errors.New("invalid token")
	//}
	//
	//if !u.Admin {
	//	return 0, errors.New("unauthorized")
	//}
	//
	//return int(userID), nil
	return 0, nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req requests.DeleteUserRequest) error {
	////adminID, err := s.CheckAdmin(ctx, r)
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//err = s.repository.DeleteUser(ctx, uint(int(req.UserID)))
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request, req requests.GetUserRequest) (domain2.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return domain2.User{}, err
	//}
	//
	//u, err := s.repository.GetUser(ctx, uint(int(req.UserID)))
	//if err != nil {
	//	return domain2.User{}, err
	//}
	//
	//return u, nil
	return domain2.User{}, nil
}

func (s Service) ProvisionDemoUser(ctx context.Context, r *http.Request, req requests.ProvisionDemoUserRequest) error {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	//cipheredPassword := s.EncryptPassword(ctx, req.Password)
	//u := domain2.User{
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
	//demoCategory := domain2.Category{
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
	//anotherDemoCategory := domain2.Category{
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
	//yetAnotherDemoCategory := domain2.Category{
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

func (s Service) List(ctx context.Context, r *http.Request) ([]domain2.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return []domain2.User{}, err
	//}
	//
	//users, err := s.repository.GetUsers(ctx)
	//if err != nil {
	//	return []domain2.User{}, err
	//}
	//
	//return users, nil
	return []domain2.User{}, nil
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

func (s Service) Activate(ctx context.Context, r requests.ActivateUserRequest) error {
	//err := s.repository.ActivateUser(ctx, r.ActivationCode)
	//if err != nil {
	//	return err
	//}

	return nil
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

func NewUserService(gr user.Repository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: gr,
	}
}
