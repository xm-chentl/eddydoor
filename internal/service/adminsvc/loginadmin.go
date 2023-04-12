package adminsvc

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/xm-chentl/eddydoor/internal/model/views"
)

type LoginAdmin struct {
	jwt.RegisteredClaims

	ID        string `json:"id"`
	Account   string `json:"account"`
	Nickname  string `json:"nickname"`
	LoginTime int64  `json:"login_time"`
	// todo: .. 角色相关
}

func (l LoginAdmin) ToResponse() views.ResponseAdmin {
	return views.ResponseAdmin{
		UserName: l.Account,
		Nickname: l.Nickname,
		Roles:    []string{"admin"},
	}
}
