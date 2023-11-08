package pkg

import (
	"errors"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func GenerateDemoTodos(categoryID, anotherCategoryID, yetAnotherCategoryID int, language string) []domain2.Todo {
	if language == "en" {
		return []domain2.Todo{
			{
				CategoryID: uint(categoryID),
				//Description: "Change Anna's diapers",
				Name:      "Diapers",
				Priority:  domain2.Priority(5),
				Recurring: true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Laundry",
				Priority:   domain2.Priority(3),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Iron clothes",
				Priority:   domain2.Priority(1),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Repair TV",
				//Description: "Need to purchase an adapter",
				Priority:  domain2.Priority(2),
				Recurring: false,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Walk Barky",
				Priority:   domain2.Priority(4),
				Recurring:  true,
			},
			{
				CategoryID: uint(categoryID),
				Name:       "Go to the veterynary",
				Priority:   domain2.Priority(4),
				Recurring:  false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Speak with Alberto",
				//Description: "Talk about the requirements for the new project",
				Priority:  domain2.Priority(3),
				Recurring: false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Speak with Laura",
				//Description: "Need to determine why the project is delayed",
				Priority:  domain2.Priority(5),
				Recurring: false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Ask for a raise",
				Priority:   domain2.Priority(3),
				Recurring:  false,
			},
			{
				CategoryID: uint(anotherCategoryID),
				Name:       "Elaborate graphs",
				//Description: "Need to elaborate some graphs for tomorrows' presentation",
				Priority:  domain2.Priority(5),
				Recurring: false,
				Active:    true,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Renew Amazon Prime",
				//Description: "Need to renew Amazon Prime before the end of the month",
				Priority:  domain2.Priority(3),
				Recurring: false,
				Active:    true,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Cancel Disney+",
				//Description: "Need to cancel Disney+ before the end of the month",
				Priority:  domain2.Priority(5),
				Recurring: false,
				Active:    false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Cancel Netflix",
				//Description: "Need to cancel Netflix before the end of the month",
				Priority:  domain2.Priority(5),
				Recurring: false,
				Active:    false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Diapers",
				Priority:   domain2.Priority(5),
				Recurring:  false,
				Active:     false,
			},
			{
				CategoryID: uint(yetAnotherCategoryID),
				Name:       "Graphic card",
				//Description: "Buy water cooled graphic card when there is a good deal",
				Priority:  domain2.Priority(1),
				Recurring: false,
				Active:    false,
			},
		}
	}

	return []domain2.Todo{
		{
			CategoryID: uint(categoryID),
			//Description: "Cambiar los pañales de Ana",
			Name:      "Pañales",
			Priority:  domain2.Priority(5),
			Recurring: true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Colada",
			Priority:   domain2.Priority(3),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Planchar la ropa",
			Priority:   domain2.Priority(1),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Arreglar la TV",
			//Description: "Comprar un adaptador",
			Priority:  domain2.Priority(2),
			Recurring: false,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Sacar a pasear a Jara",
			Priority:   domain2.Priority(4),
			Recurring:  true,
		},
		{
			CategoryID: uint(categoryID),
			Name:       "Ir al veterinario",
			Priority:   domain2.Priority(4),
			Recurring:  false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Hablar con Laura",
			//Description: "Hablar respecto a los requisitos del nuevo proyecto",
			Priority:  domain2.Priority(3),
			Recurring: false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Hablar con Alberto",
			//Description: "Determinar por qué el proyecto va con retraso",
			Priority:  domain2.Priority(5),
			Recurring: false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Pedir un aumento",
			Priority:   domain2.Priority(3),
			Recurring:  false,
		},
		{
			CategoryID: uint(anotherCategoryID),
			Name:       "Elaborar gráficos",
			//Description: "Elaborar gráficos para la presentación de mañana",
			Priority:  domain2.Priority(5),
			Recurring: false,
			Active:    true,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Renovar Amazon Prime",
			//Description: "Renovar Amazon Prime antes del final del mes",
			Priority:  domain2.Priority(3),
			Recurring: false,
			Active:    true,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Cancelar Disney+",
			//Description: "Cancelar Disney+ antes del final del mes",
			Priority:  domain2.Priority(5),
			Recurring: false,
			Active:    false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Cancelar Netflix",
			//Description: "Cancelar Netflix antes del final del mes",
			Priority:  domain2.Priority(5),
			Recurring: false,
			Active:    false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Pañales",
			Priority:   domain2.Priority(5),
			Recurring:  false,
			Active:     false,
		},
		{
			CategoryID: uint(yetAnotherCategoryID),
			Name:       "Tarjeta gráfica",
			//Description: "Comprar tarjeta gráfica con refrigeración líquida cuando haya una buena oferta",
			Priority:  domain2.Priority(1),
			Recurring: false,
			Active:    false,
		},
	}
}

//func FilterUsers(users []domain.User) []domain.FilteredUser {
//	filteredUsers := make([]domain.FilteredUser, 0, len(users))
//	for _, user := range users {
//		filteredUsers = append(filteredUsers, FilterUser(user))
//	}
//	return filteredUsers
//}
//
//func FilterUser(user domain.User) domain.FilteredUser {
//	return domain.FilteredUser{
//		ID:               user.ID,
//		Email:            user.Email,
//		Name:             user.Name,
//		RegistrationDate: user.RegistrationDate,
//	}
//}

func SendEmail(e domain2.Email) error {
	from := mail.NewEmail(FromName, FromEmail)
	subject := e.Subject
	to := mail.NewEmail(e.Recipient, e.To)
	message := mail.NewSingleEmail(from, subject, to, e.Body+time.Now().Format("2006-01-02 15:04:05"), e.Body)
	client := sendgrid.NewSendClient(SendGridApiKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		return errors.New("error sending email")
	}
	return nil
}

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}

func GetOrigin() string {
	if Environment == "local" {
		return DapsLocalUrl
	}
	return DapsProductionUrl
}

//func GenerateRemindTodosHTMLContent(u domain2.User, rs []domain2.RemindSummary) (domain2.Email, error) {
//	e := domain2.Email{}
//	reminders := make(map[string][]domain2.RemindSummary)
//	for _, r := range rs {
//		if _, ok := reminders[r.CategoryName]; !ok {
//			reminders[r.CategoryName] = []domain2.RemindSummary{r}
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
//		Reminders map[string][]domain2.RemindSummary
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
