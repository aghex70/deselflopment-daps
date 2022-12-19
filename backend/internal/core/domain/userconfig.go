package domain

type UserConfig struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	AutoSuggest bool   `json:"auto_suggest"`
	Language    string `json:"language"`
}

type Profile struct {
	Email string `json:"email"`
	UserConfig
}
