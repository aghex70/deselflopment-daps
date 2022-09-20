package todo

import (
	"github.com/aghex70/daps/internal/core/ports"
	"log"
)

type RestHandler struct {
	categoryService ports.CategoryServicer
	logger          *log.Logger
}
