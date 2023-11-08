package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/category"
	"github.com/aghex70/daps/internal/ports/requests/category"
	"log"
	"net/http"
)

type Service struct {
	logger     *log.Logger
	repository category.Repository
}

func (s Service) Create(ctx context.Context, r *http.Request, req requests.CreateCategoryRequest) error {
	//userID, _ := server.RetrieveJWTClaims(r, req)
	//err := s.ValidateCreation(ctx, req.Name, userID)
	//if err != nil {
	//	return err
	//}
	//u := domain2.User{ID: userID}
	//cat := domain2.Category{
	//	OwnerID: userID,
	//	Name:    req.Name,
	//	//Description: req.Description,
	//	Custom:     true,
	//	Notifiable: req.Notifiable,
	//	Users:      &[]domain2.User{u},
	//}
	//_, err = s.repository.Create(ctx, cat)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s Service) Update(ctx context.Context, r *http.Request, req requests.UpdateCategoryRequest) error {
	//userID, _ := server.RetrieveJWTClaims(r, req)
	//
	//switch {
	//case req.Shared == nil:
	//	err := s.ValidateModification(ctx, uint(int(req.CategoryID)), uint(int(userID)))
	//	if err != nil {
	//		return err
	//	}
	//	cat := domain.Category{
	//		ID:          req.CategoryID,
	//		Description: req.Description,
	//		Notifiable:  req.Notifiable,
	//		Name:        req.Name,
	//	}
	//	err = s.repository.UpdateCategory(ctx, cat)
	//	if err != nil {
	//		return err
	//	}
	//
	//case *req.Shared:
	//	err := s.ValidateShare(ctx, uint(int(req.CategoryID)), uint(int(userID)))
	//	if err != nil {
	//		return err
	//	}
	//	cat := domain.Category{
	//		ID:     req.CategoryID,
	//		Shared: *req.Shared,
	//	}
	//	err = s.repository.Share(ctx, cat, req.Email)
	//	if err != nil {
	//		return err
	//	}
	//
	//case !*req.Shared:
	//	err := s.ValidateUnshare(ctx, uint(int(req.CategoryID)), uint(int(userID)))
	//	if err != nil {
	//		return err
	//	}
	//	cat := domain.Category{
	//		ID:     req.CategoryID,
	//		Shared: *req.Shared,
	//	}
	//	err = s.repository.Unshare(ctx, cat, int(userID))
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request, req requests.GetCategoryRequest) (domain.Category, error) {
	//userID, _ := server.RetrieveJWTClaims(r, req)
	//err := s.ValidateRetrieval(ctx, uint(int(req.CategoryID)), uint(int(userID)))
	//if err != nil {
	//	return domain2.Category{}, err
	//}
	//cat, err := s.repository.Get(ctx, uint(int(req.CategoryID)))
	//if err != nil {
	//	return domain2.Category{}, err
	//}
	//return cat, nil
	return domain.Category{}, nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req requests.DeleteCategoryRequest) error {
	//userID, _ := server.RetrieveJWTClaims(r, req)
	//err := s.ValidateRemoval(ctx, uint(int(req.CategoryID)), uint(int(userID)))
	//if err != nil {
	//	return err
	//}
	//err = s.repository.DeleteCategory(ctx, req.CategoryID, userID)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s Service) List(ctx context.Context, r *http.Request) ([]domain.Category, error) {
	//userID, _ := server.RetrieveJWTClaims(r, nil)
	//categories, err := s.repository.GetCategories(ctx, userID)
	//if err != nil {
	//	return []domain.Category{}, err
	//}
	//return categories, nil
	return []domain.Category{}, nil
}

func NewCategoryService(r category.Repository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
