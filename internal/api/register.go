package api

import (
	"github.com/xm-chentl/eddydoor/internal/api/app"
	"github.com/xm-chentl/eddydoor/internal/api/user"
	"github.com/xm-chentl/eddydoor/utils/handlers"
)

func init() {
	handlers.Register([]handlers.HandlerItem{
		{"/app/get": &app.GetAPI{}},
		{"/user/get-login-sms-code": &user.GetLoginSMSCodeAPI{}},
		{"/user/get-register-sms-code": &user.GetRegisterSMSCodeAPI{}},
		{"/user/register": &user.RegisterAPI{}},
		{"/user/login": &user.LoginAPI{}},
		{"/user/login-by-phone": &user.LoginByPhoneAPI{}},
		{"/user/refresh-token": &user.RefreshTokenAPI{}},
	}...)
}
