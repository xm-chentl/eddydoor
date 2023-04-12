package contract

import "github.com/xm-chentl/eddydoor/internal/service/adminsvc"

type IAdminSession interface {
	GetAdmin() adminsvc.LoginAdmin
}
