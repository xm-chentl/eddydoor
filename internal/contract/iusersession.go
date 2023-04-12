package contract

import "github.com/xm-chentl/eddydoor/internal/service/usersvc"

type IUserSession interface {
	GetUser() usersvc.LoginUser
}
