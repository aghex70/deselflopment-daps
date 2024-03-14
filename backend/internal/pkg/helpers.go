package pkg

import (
	"github.com/aghex70/daps/internal/ports/domain"
	"net/http"
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

func FilterUsers(users []domain.User) []domain.FilteredUser {
	filteredUsers := make([]domain.FilteredUser, 0, len(users))
	for _, user := range users {
		filteredUsers = append(filteredUsers, FilterUser(user))
	}
	return filteredUsers
}

func FilterUser(user domain.User) domain.FilteredUser {
	return domain.FilteredUser{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Admin:     user.Admin,
		Active:    user.Active,
	}
}

func FilterProfile(user domain.User) domain.Profile {
	return domain.Profile{
		Email:       user.Email,
		AutoSuggest: user.AutoSuggest,
		Language:    user.Language,
	}
}

func FilterCategories(categories []domain.Category) []domain.FilteredCategory {
	filteredCategories := make([]domain.FilteredCategory, 0, len(categories))
	for _, category := range categories {
		filteredCategories = append(filteredCategories, FilterCategory(category))
	}
	return filteredCategories
}

func FilterCategory(category domain.Category) domain.FilteredCategory {
	return domain.FilteredCategory{
		ID:          category.ID,
		CreatedAt:   category.CreatedAt,
		Name:        category.Name,
		Description: category.Description,
		Shared:      category.Shared,
		Notifiable:  category.Notifiable,
		Custom:      category.Custom,
	}
}

func SetCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", GetOrigin())
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
}
func GetOrigin() string {
	if Environment == "local" {
		return DapsLocalUrl
	}
	return DapsProductionUrl
}

//func GenerateRemindTodosHTMLContent(u domain.User, rs []domain.RemindSummary) (domain.Email, error) {
//	e := domain.Email{}
//	reminders := make(map[string][]domain.RemindSummary)
//	for _, r := range rs {
//		if _, ok := reminders[r.CategoryName]; !ok {
//			reminders[r.CategoryName] = []domain.RemindSummary{r}
//		} else {
//			reminders[r.CategoryName] = append(reminders[r.CategoryName], r)
//		}
//	}
//
//	var tpl bytes.Buffer
//	t := template.Must(template.New("emailTemplate").Parse(`
//		<html>
//			<body>
//				<h2>Hola {{.Name}},</h2>
//				<h2>{{.Header}}</h2>
//				{{range $category, $todos := .Reminders}}
//				<h2>- {{$category}}</h3>
//				<ul>
//					{{range $i, $todo := $todos}}
//					{{if eq $todo.TodoPriority 1}}
//					<li style="color:gray">
//						{{$todo.TodoName}}
//						{{if $todo.TodoDescription}}
//						({{$todo.TodoDescription}})
//						{{end}}
//						{{if $todo.TodoLink}}
//						<a href="{{$todo.TodoLink}}">Link</a>
//						{{end}}
//					</li>
//					{{else if eq $todo.TodoPriority 2}}
//					<li style="color:blue">
//						{{$todo.TodoName}}
//						{{if $todo.TodoDescription}}
//						({{$todo.TodoDescription}})
//						{{end}}
//						{{if $todo.TodoLink}}
//						<a href="{{$todo.TodoLink}}">Link</a>
//						{{end}}
//					</li>
//					{{else if eq $todo.TodoPriority 3}}
//					<li style="color:green">
//						{{$todo.TodoName}}
//						{{if $todo.TodoDescription}}
//						({{$todo.TodoDescription}})
//						{{end}}
//						{{if $todo.TodoLink}}
//						<a href="{{$todo.TodoLink}}">Link</a>
//						{{end}}
//					</li>
//					{{else if eq $todo.TodoPriority 4}}
//					<li style="color:yellow">
//						{{$todo.TodoName}}
//						{{if $todo.TodoDescription}}
//						({{$todo.TodoDescription}})
//						{{end}}
//						{{if $todo.TodoLink}}
//						<a href="{{$todo.TodoLink}}">Link</a>
//						{{end}}
//					</li>
//					{{else if eq $todo.TodoPriority 5}}
//					<li style="color:red">
//						{{$todo.TodoName}}
//						{{if $todo.TodoDescription}}
//						({{$todo.TodoDescription}})
//						{{end}}
//						{{if $todo.TodoLink}}
//						<a href="{{$todo.TodoLink}}">Link</a>
//						{{end}}
//					</li>
//					{{end}}
//					{{end}}
//				</ul>
//				{{end}}
//			</body>
//		</html>
//	`))
//
//	data := struct {
//		Name      string
//		Header    string
//		Reminders map[string][]domain.RemindSummary
//	}{
//		Name:      u.Name,
//		Header:    "Aquí tienes (algunas) de tus tareas pendientes a fecha de " + time.Now().Format("02/01/2006"+":"),
//		Reminders: reminders,
//	}
//
//	err := t.Execute(&tpl, data)
//	if err != nil {
//		return e, err
//	}
//
//	e.Body = tpl.String()
//	return e, nil
//}
