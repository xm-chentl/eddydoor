package views

type ResponseAdmin struct {
	UserName     string   `json:"username"`
	Nickname     string   `json:"name"`
	Avatar       string   `json:"avatar"`
	Introduction string   `json:"introduction"`
	EMail        string   `json:"email"`
	Phone        string   `json:"phone"`
	Roles        []string `json:"roles"`
}
