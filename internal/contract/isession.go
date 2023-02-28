package contract

import "github.com/xm-chentl/eddydoor/internal/service/usersvc"

type ISession interface {
	GetUser() usersvc.LoginUser
}
