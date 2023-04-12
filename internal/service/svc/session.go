package svc

import (
	"github.com/xm-chentl/eddydoor/internal/service/adminsvc"
	"github.com/xm-chentl/eddydoor/internal/service/usersvc"
)

type UserSession struct {
	User usersvc.LoginUser
}

func (s UserSession) GetUser() usersvc.LoginUser {
	return s.User
}

type AdminSession struct {
	Admin adminsvc.LoginAdmin
}

func (s AdminSession) GetAdmin() adminsvc.LoginAdmin {
	return s.Admin
}
