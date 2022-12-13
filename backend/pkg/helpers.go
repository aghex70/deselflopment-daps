package pkg

import "github.com/aghex70/daps/internal/core/domain"


func GenerateDemoTodos(categoryId, anotherCategoryId int, language string) []domain.Todo {
	if language == "en" {
		return []domain.Todo{
			{
				Category:    categoryId,
				Description: "Change Anna's diapers",
				Name:        "Diapers",
				Priority:    domain.Priority(5),
				Recurring:   true,
			},
			{
				Category:  categoryId,
				Name:      "Laundry",
				Priority:  domain.Priority(3),
				Recurring: true,
			},
			{
				Category:  categoryId,
				Name:      "Iron clothes",
				Priority:  domain.Priority(1),
				Recurring: true,
			},
			{
				Category:    categoryId,
				Name:        "Repair TV",
				Description: "Need to purchase an adapter",
				Priority:    domain.Priority(2),
				Recurring:   false,
			},
			{
				Category:  categoryId,
				Name:      "Walk Barky",
				Priority:  domain.Priority(4),
				Recurring: true,
			},
			{
				Category:  categoryId,
				Name:      "Go to the veterynary",
				Priority:  domain.Priority(4),
				Recurring: false,
			},
			{
				Category:    anotherCategoryId,
				Name:        "Speak with Alberto",
				Description: "Talk about the requirements for the new project",
				Priority:    domain.Priority(3),
				Recurring:   false,
			},
			{
				Category:    anotherCategoryId,
				Name:        "Speak with Laura",
				Description: "Need to determine why the project is delayed",
				Priority:    domain.Priority(5),
				Recurring:   false,
			},
			{
				Category:  anotherCategoryId,
				Name:      "Ask for a raise",
				Priority:  domain.Priority(3),
				Recurring: false,
			},
			{
				Category:    anotherCategoryId,
				Name:        "Elaborate graphs",
				Description: "Need to elaborate some graphs for tomorrows' presentation",
				Priority:    domain.Priority(5),
				Recurring:   false,
			},
		}
	}

	return []domain.Todo{
		{
			Category:    categoryId,
			Description: "Cambiar los pañales de Ana",
			Name:        "Pañales",
			Priority:    domain.Priority(5),
			Recurring:   true,
		},
		{
			Category:  categoryId,
			Name:      "Colada",
			Priority:  domain.Priority(3),
			Recurring: true,
		},
		{
			Category:  categoryId,
			Name:      "Planchar la ropa",
			Priority:  domain.Priority(1),
			Recurring: true,
		},
		{
			Category:    categoryId,
			Name:        "Arreglar la TV",
			Description: "Comprar un adaptador",
			Priority:    domain.Priority(2),
			Recurring:   false,
		},
		{
			Category:  categoryId,
			Name:      "Sacar a pasear a Jara",
			Priority:  domain.Priority(4),
			Recurring: true,
		},
		{
			Category:  categoryId,
			Name:      "Ir al veterinario",
			Priority:  domain.Priority(4),
			Recurring: false,
		},
		{
			Category:    anotherCategoryId,
			Name:        "Hablar con Laura",
			Description: "Hablar respecto a los requisitos del nuevo proyecto",
			Priority:    domain.Priority(3),
			Recurring:   false,
		},
		{
			Category:    anotherCategoryId,
			Name:        "Hablar con Alberto",
			Description: "Determinar por qué el proyecto va con retraso",
			Priority:    domain.Priority(5),
			Recurring:   false,
		},
		{
			Category:  anotherCategoryId,
			Name:      "Pedir un aumento",
			Priority:  domain.Priority(3),
			Recurring: false,
		},
		{
			Category:    anotherCategoryId,
			Name:        "Elaborar gráficos",
			Description: "Elaborar gráficos para la presentación de mañana",
			Priority:    domain.Priority(5),
			Recurring:   false,
		},
	}
}
