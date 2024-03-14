package email

import (
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
)

func GenerateActivationEmail(to, recipient string, u domain.User) domain.Email {
	e := domain.Email{
		Subject:   "📣 DAPS - Activate your account 📣",
		Body:      "In order to complete your registration, please click on the following link: " + pkg.ActivationCodeLink + u.ActivationCode,
		From:      pkg.FromEmail,
		Source:    pkg.ProjectName,
		To:        to,
		Recipient: recipient,
		UserID:    u.ID,
	}
	return e
}

func GenerateResetPasswordEmail(u domain.User) domain.Email {
	e := domain.Email{
		Subject:   "📣 DAPS - Reset your password 📣",
		Body:      "In order to reset your password, please follow this link: " + pkg.ResetPasswordLink + u.ResetPasswordCode,
		From:      pkg.FromEmail,
		Source:    pkg.ProjectName,
		To:        u.Name,
		Recipient: u.Email,
		UserID:    u.ID,
	}
	return e
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
