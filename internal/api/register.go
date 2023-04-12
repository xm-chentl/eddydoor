package api

import (
	"github.com/xm-chentl/eddydoor/internal/api/admin"
	"github.com/xm-chentl/eddydoor/internal/api/app"
	"github.com/xm-chentl/eddydoor/internal/api/config"
	"github.com/xm-chentl/eddydoor/internal/api/user"
	"github.com/xm-chentl/eddydoor/plugin/muxsvr"
)

func init() {
	muxsvr.RegisterHandlers([]muxsvr.Handler{
		{"/admin/login": &admin.LoginAPI{}},
		{"/admin/info": &admin.InfoAPI{}},
		{"/app/get": &app.GetAPI{}},
		{"/config/modify": &config.ModifyAPI{}},
		{"/user/get-login-sms-code": &user.GetLoginSMSCodeAPI{}},
		{"/user/get-register-sms-code": &user.GetRegisterSMSCodeAPI{}},
		{"/user/register": &user.RegisterAPI{}},
		{"/user/login-by-phone": &user.LoginByPhoneAPI{}},
		{"/user/refresh-token": &user.RefreshTokenAPI{}},
		{"/user/list": &user.ListAPI{}},
	}...)
}
