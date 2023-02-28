package ginex

import "github.com/xm-chentl/eddydoor/internal/service/usersvc"

type Session struct {
	User usersvc.LoginUser
}

func (s Session) GetUser() usersvc.LoginUser {
	return s.User
}
