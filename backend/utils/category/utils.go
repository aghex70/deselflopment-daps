package category

import (
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
)

func GenerateDemoTodos(categoryID, anotherCategoryID, yetAnotherCategoryID int, language string) []domain.Todo {
	if language == "en" {
		return []domain.Todo{
			{
				CategoryID: uint(categoryID),
				//Description: "Change Anna's diapers",
				Name:      "Diapers",
				Priority:  domain.Priority(5),
				Recurring: true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Laundry",
				Priority:   domain.Priority(3),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Iron clothes",
				Priority:   domain.Priority(1),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Repair TV",
				//Description: "Need to purchase an adapter",
				Priority:  domain.Priority(2),
				Recurring: false,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Walk Barky",
				Priority:   domain.Priority(4),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Go to the veterynary",
				Priority:   domain.Priority(4),
				Recurring:  false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Speak with Alberto",
				//Description: "Talk about the requirements for the new project",
				Priority:  domain.Priority(3),
				Recurring: false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Speak with Laura",
				//Description: "Need to determine why the project is delayed",
				Priority:  domain.Priority(5),
				Recurring: false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Ask for a raise",
				Priority:   domain.Priority(3),
				Recurring:  false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Elaborate graphs",
				//Description: "Need to elaborate some graphs for tomorrows' presentation",
				Priority:  domain.Priority(5),
				Recurring: false,
				Active:    true,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Renew Amazon Prime",
				//Description: "Need to renew Amazon Prime before the end of the month",
				Priority:  domain.Priority(3),
				Recurring: false,
				Active:    true,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Cancel Disney+",
				//Description: "Need to cancel Disney+ before the end of the month",
				Priority:  domain.Priority(5),
				Recurring: false,
				Active:    false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Cancel Netflix",
				//Description: "Need to cancel Netflix before the end of the month",
				Priority:  domain.Priority(5),
				Recurring: false,
				Active:    false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Diapers",
				Priority:   domain.Priority(5),
				Recurring:  false,
				Active:     false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Graphic card",
				//Description: "Buy water cooled graphic card when there is a good deal",
				Priority:  domain.Priority(1),
				Recurring: false,
				Active:    false,
			},
		}
	}

	return []domain.Todo{
		{
			CategoryID: uint(categoryID),
			//Description: "Cambiar los pañales de Ana",
			Name:      "Pañales",
			Priority:  domain.Priority(5),
			Recurring: true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Colada",
			Priority:   domain.Priority(3),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Planchar la ropa",
			Priority:   domain.Priority(1),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Arreglar la TV",
			//Description: "Comprar un adaptador",
			Priority:  domain.Priority(2),
			Recurring: false,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Sacar a pasear a Jara",
			Priority:   domain.Priority(4),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Ir al veterinario",
			Priority:   domain.Priority(4),
			Recurring:  false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Hablar con Laura",
			//Description: "Hablar respecto a los requisitos del nuevo proyecto",
			Priority:  domain.Priority(3),
			Recurring: false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Hablar con Alberto",
			//Description: "Determinar por qué el proyecto va con retraso",
			Priority:  domain.Priority(5),
			Recurring: false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Pedir un aumento",
			Priority:   domain.Priority(3),
			Recurring:  false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Elaborar gráficos",
			//Description: "Elaborar gráficos para la presentación de mañana",
			Priority:  domain.Priority(5),
			Recurring: false,
			Active:    true,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Renovar Amazon Prime",
			//Description: "Renovar Amazon Prime antes del final del mes",
			Priority:  domain.Priority(3),
			Recurring: false,
			Active:    true,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Cancelar Disney+",
			//Description: "Cancelar Disney+ antes del final del mes",
			Priority:  domain.Priority(5),
			Recurring: false,
			Active:    false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Cancelar Netflix",
			//Description: "Cancelar Netflix antes del final del mes",
			Priority:  domain.Priority(5),
			Recurring: false,
			Active:    false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Pañales",
			Priority:   domain.Priority(5),
			Recurring:  false,
			Active:     false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Tarjeta gráfica",
			//Description: "Comprar tarjeta gráfica con refrigeración líquida cuando haya una buena oferta",
			Priority:  domain.Priority(1),
			Recurring: false,
			Active:    false,
		},
	}
}

func IsCategoryOwner(ownerID, userID uint) bool {
	return ownerID == userID
}

func CanRetrieveCategory(cs []domain.Category, id uint) (domain.Category, error) {
	for _, c := range cs {
		if c.ID == id {
			return c, nil
		}
	}
	return domain.Category{}, pkg.UnauthorizedError
}

func CanEditCategory(cs []domain.Category, id uint) (domain.Category, error) {
	for _, c := range cs {
		if c.ID == id {
			return c, nil
		}
	}
	return domain.Category{}, pkg.UnauthorizedError
}
