package category

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/server"
	"log"
	"net/http"
	"strconv"
)

type CategoryService struct {
	logger                 *log.Logger
	categoryRepository     *category.CategoryGormRepository
	relationshipRepository *relationship.RelationshipGormRepository
	todoRepository         *todo.TodoGormRepository
}

func (s CategoryService) Create(ctx context.Context, r *http.Request, req ports.CreateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateCreation(ctx, req.Name, int(userId))
	if err != nil {
		return err
	}
	u := domain.User{ID: int(userId)}
	cat := domain.Category{
		OwnerID:           int(userId),
		Description:       req.Description,
		Custom:            true,
		Name:              req.Name,
		InternationalName: req.InternationalName,
		Users:             []domain.User{u},
	}
	c, err := s.categoryRepository.Create(ctx, cat, int(userId))
	if err != nil {
		return err
	}

	fmt.Println(c)
	//nntd := domain.Todo{
	//	Category:    c.ID,
	//	Description: "This is a default todo",
	//	Name:        "Default todo",
	//	Link:        "Default URL",
	//	Priority:    domain.Priority(2),
	//	Recurring:   false,
	//}
	//
	//fmt.Printf("\n\n%+v", c)
	//fmt.Printf("\n\n%+v", nntd)
	//err = s.todoRepository.Create(ctx, nntd)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s CategoryService) Update(ctx context.Context, r *http.Request, req ports.UpdateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)

	if req.Shared == nil {
		err := s.ValidateModification(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:                int(req.CategoryId),
			Description:       req.Description,
			Name:              req.Name,
			InternationalName: req.InternationalName,
		}
		err = s.categoryRepository.Update(ctx, cat)
		if err != nil {
			return err
		}
	} else if *req.Shared == true {
		err := s.ValidateShare(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:     int(req.CategoryId),
			Shared: req.Shared,
		}
		err = s.categoryRepository.Share(ctx, cat, req.Email)
		if err != nil {
			return err
		}

	} else if *req.Shared == false {
		err := s.ValidateUnshare(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:     int(req.CategoryId),
			Shared: req.Shared,
		}
		err = s.categoryRepository.Unshare(ctx, cat, int(userId))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s CategoryService) Get(ctx context.Context, r *http.Request, req ports.GetCategoryRequest) (domain.Category, error) {
	fmt.Println("99999999999999999999")
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateRetrieval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		fmt.Println("%+v", err)
		fmt.Println("qweqweqweqweqweqweqwe")
		return domain.Category{}, err
	}
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	td, err := s.categoryRepository.GetById(ctx, int(req.CategoryId))
	if err != nil {
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		return domain.Category{}, err
	}
	return td, nil
}

func (s CategoryService) Delete(ctx context.Context, r *http.Request, req ports.DeleteCategoryRequest) error {
	fmt.Printf("\n\nrequest ------>: %+v", r)
	q := r.URL.Query()
	bod := r.Body
	fmt.Printf("\n\nqparams ------>: %+v", q)
	fmt.Printf("\n\nbody ------>: %+v", bod)
	payload := ports.ListTodosRequest{}
	//err := handlers.ValidateRequest(r, &payload)
	//if err != nil {
	//	handlers.ThrowError(err, http.StatusBadRequest, w)
	//	return
	//}

	categoryId, err := strconv.Atoi(q.Get("category_id"))
	payload.Category = categoryId
	//todos, err := h.toDoService.List(nil, r, payload)

	userId, _ := server.RetrieveJWTClaims(r, req)
	err = s.ValidateRemoval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	fmt.Println("Trying to deleeeeeeeeeeeeeeeeeeeeeete")
	err = s.categoryRepository.Delete(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		fmt.Println("Error deleting @@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		return err
	}
	return nil
}

func (s CategoryService) List(ctx context.Context, r *http.Request) ([]domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	todos, err := s.categoryRepository.List(ctx, int(userId))
	if err != nil {
		return []domain.Category{}, err
	}
	return todos, nil
}

func NewCategoryService(cr *category.CategoryGormRepository, rr *relationship.RelationshipGormRepository, logger *log.Logger) CategoryService {
	return CategoryService{
		logger:                 logger,
		categoryRepository:     cr,
		relationshipRepository: rr,
	}
}
