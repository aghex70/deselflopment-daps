package common

import (
	"github.com/satori/go.uuid"
)

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
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

//func SendEmail(e domain.Email) error {
//	from := mail.NewEmail(FromName, FromEmail)
//	subject := e.Subject
//	to := mail.NewEmail(e.Recipient, e.To)
//	message := mail.NewSingleEmail(from, subject, to, e.Body+time.Now().Format("2006-01-02 15:04:05"), e.Body)
//	client := sendgrid.NewSendClient(SendGridApiKey)
//	response, err := client.Send(message)
//	if err != nil {
//		return err
//	}
//	if response.StatusCode != 202 {
//		return errors.New("error sending email")
//	}
//	return nil
//}
//
//func GenerateUUID() string {
//	u := uuid.NewV4()
//	return u.String()
//}
//
//func GetOrigin() string {
//	if Environment == "local" {
//		return DapsLocalUrl
//	}
//	return DapsProductionUrl
//}

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
