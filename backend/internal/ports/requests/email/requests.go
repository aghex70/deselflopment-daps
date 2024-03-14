package requests

type SendEmailRequest struct {
	To        string `json:"to"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	UserId    int64  `json:"user_id"`
}
